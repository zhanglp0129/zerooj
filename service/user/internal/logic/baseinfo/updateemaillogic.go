package baseinfologic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"time"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户邮箱，有7天冷却期
func (l *UpdateEmailLogic) UpdateEmail(in *user.UpdateEmailReq) (*user.UpdateEmailResp, error) {
	// 判断是否处于冷却期
	rdb := l.svcCtx.RDB
	coolPeriodKey := fmt.Sprintf("/frequency_limit/user/update_email/week/%d", in.Id)
	err := common.InCoolPeriod(rdb, coolPeriodKey)
	if err != nil {
		return nil, err
	}

	// 先获取旧邮箱
	db := l.svcCtx.DB
	u := models.User{}
	err = db.Select("password", "email").Take(&u, in.Id).Error
	if err != nil {
		return nil, err
	}

	// 判断新旧邮箱是否相同
	if in.Email == u.Email {
		return nil, constant.NewEqualsOldError{Thing: "邮箱"}
	}

	oldKey := fmt.Sprintf("/mail_check_code/user/update_email/old/%s", u.Email)
	newKey := fmt.Sprintf("/mail_check_code/user/update_email/new/%s", in.Email)
	if in.OldPassword != "" {
		// 校验密码，只需与 旧邮箱验证码 校验一个
		oldPasswordEncrypt, err := utils.PasswordEncrypt(in.OldPassword)
		if err != nil {
			return nil, err
		}
		if oldPasswordEncrypt != u.Password {
			return nil, constant.InputDataError{Thing: "旧密码"}
		}
	} else {
		// 校验旧邮箱验证码
		trueOldCheckCode, err := rdb.Get(context.Background(), oldKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		} else if errors.Is(err, redis.Nil) || trueOldCheckCode != in.OldEmailCheckCode {
			return nil, constant.InputDataError{Thing: "旧邮箱验证码"}
		}
	}

	// 校验新邮箱验证码
	trueONewCheckCode, err := rdb.Get(context.Background(), newKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	} else if errors.Is(err, redis.Nil) || trueONewCheckCode != in.EmailCheckCode {
		return nil, constant.InputDataError{Thing: "新邮箱验证码"}
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 修改邮箱
		err = tx.Model(&models.User{}).Where("id = ?", in.Id).Update("email", in.Email).Error
		if err != nil {
			return err
		}

		getBaseInfoKey := fmt.Sprintf("/cache/user/get_base_info/%d", in.Id)
		err = rdb.Watch(context.Background(), func(rtx *redis.Tx) error {
			pipe := rtx.Pipeline()
			// 删除缓存
			redis_cache.DeleteCacheToPipe(pipe, getBaseInfoKey)
			// 设置冷却期
			period := 7 * 24 * time.Hour
			common.AddCoolPeriodToPipe(pipe, coolPeriodKey, period)

			// 执行
			_, err = pipe.Exec(context.Background())

			return err
		}, coolPeriodKey, getBaseInfoKey)
		return err
	})
	if err != nil {
		return nil, err
	}

	// 邮箱验证码使用完成
	go common.FinishMailCheckCode(rdb, oldKey)
	go common.FinishMailCheckCode(rdb, newKey)

	return &user.UpdateEmailResp{
		OldEmail: u.Email,
	}, nil
}
