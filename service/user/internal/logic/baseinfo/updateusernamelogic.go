package baseinfologic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"time"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/models"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsernameLogic {
	return &UpdateUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户名，有7天冷却期
func (l *UpdateUsernameLogic) UpdateUsername(in *user.UpdateUsernameReq) (*user.UpdateUsernameResp, error) {
	// 判断是否处于冷却期
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("/frequency_limit/user/update_username/week/%d", in.Id)
	err := common.InCoolPeriod(rdb, key)
	if err != nil {
		return nil, err
	}

	// 先获取旧用户名
	db := l.svcCtx.DB
	u := models.User{}
	err = db.Select("username").Where("id = ?", in.Id).Take(&u).Error
	if err != nil {
		return nil, err
	}
	oldUsername := u.Username
	if in.Username == oldUsername {
		return nil, constant.NewUsernameEqualsOldError
	}

	// 修改数据
	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&u).Where("id = ?", in.Id).Update("username", in.Username).Error
		if err != nil {
			return err
		}

		// 删除缓存
		getBaseInfoKey := fmt.Sprintf("/cache/user/get_base_info/%d", in.Id)
		searchByUsernameKey := fmt.Sprintf(" /cache/user/search_by_username/%s", oldUsername)
		err = rdb.Watch(context.Background(), func(rtx *redis.Tx) error {
			pipe := rtx.Pipeline()
			redis_cache.DeleteCacheToPipe(pipe, getBaseInfoKey)
			redis_cache.DeleteCacheToPipe(pipe, searchByUsernameKey)
			// 设置冷却期，7天
			period := 7 * 24 * time.Hour
			common.AddCoolPeriodToPipe(pipe, key, period)

			// 执行
			_, err = pipe.Exec(context.Background())
			return err
		}, key, getBaseInfoKey, searchByUsernameKey)

		return err
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdateUsernameResp{
		OldUsername: oldUsername,
	}, nil
}
