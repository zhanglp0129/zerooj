package baseinfologic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseInfoLogic {
	return &UpdateBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户基本信息，字段为空表示不修改
func (l *UpdateBaseInfoLogic) UpdateBaseInfo(in *user.UpdateBaseInfoReq) (*user.UpdateBaseInfoResp, error) {
	rdb := l.svcCtx.RDB
	db := l.svcCtx.DB
	var u models.User

	// 先获取旧用户名
	err := db.Select("username").Where("id = ?", in.Id).Take(&u).Error
	if err != nil {
		return nil, err
	}
	oldUsername := u.Username

	// 构造修改的数据模型
	encryptedPassword, err := utils.PasswordEncrypt(in.Password)
	if err != nil {
		return nil, err
	}
	u = models.User{
		Username: in.Username,
		Password: encryptedPassword,
		Email:    in.Email,
	}
	u.ID = in.Id

	// 修改数据
	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Updates(&u).Error
		if err != nil {
			return err
		}

		// 删除缓存
		userBaseInfoKey := fmt.Sprintf("cache:user_base_info:%d", in.Id)
		searchByUsernameKey := fmt.Sprintf("cache:search_by_username:%s", oldUsername)
		err = rdb.Watch(context.Background(), func(rtx *redis.Tx) error {
			pipe := rtx.Pipeline()
			redis_cache.DeleteCacheToPipe(pipe, userBaseInfoKey)
			redis_cache.DeleteCacheToPipe(pipe, searchByUsernameKey)

			// 执行
			_, err = pipe.Exec(context.Background())
			return err
		}, userBaseInfoKey, searchByUsernameKey)

		return err
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdateBaseInfoResp{}, nil
}
