package followlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"time"
	"zerooj/common/constant"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/models"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注其他用户
func (l *FollowUserLogic) FollowUser(in *user.FollowUserReq) (*user.FollowUserResp, error) {
	// 不能关注自己
	if in.FollowerId == in.FollowedId {
		return nil, errors.New(constant.FollowYourselfError)
	}

	db := l.svcCtx.DB
	follow := models.Follow{
		FollowerID: in.FollowerId,
		FollowedID: in.FollowedId,
		FollowAt:   time.Now(),
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// 添加记录
		result := tx.Create(&follow)
		rows := result.RowsAffected
		err := result.Error
		if err != nil {
			return err
		}

		if rows > 0 {
			// 修改缓存
			rdb := l.svcCtx.RDB
			followingKey := fmt.Sprintf("/cache/user/get_followings/count/%d", in.FollowerId)
			fanKey := fmt.Sprintf("/cache/user/get_fans/count/%d", in.FollowedId)
			err = rdb.Watch(context.Background(), func(rtx *redis.Tx) error {
				pipe := rtx.Pipeline()
				redis_cache.CacheIncrByToPipe(pipe, followingKey, rows)
				redis_cache.CacheIncrByToPipe(pipe, fanKey, rows)
				_, err = pipe.Exec(context.Background())
				return err
			}, followingKey, fanKey)
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &user.FollowUserResp{}, nil
}
