package followlogic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowUserLogic {
	return &UnfollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnfollowUserLogic) UnfollowUser(in *user.UnfollowUserReq) (*user.UnfollowUserResp, error) {
	db := l.svcCtx.DB
	follow := models.Follow{
		FollowerID: in.FollowerId,
		FollowedID: in.FollowedId,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// 删除记录
		result := tx.Delete(&follow)
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
				redis_cache.CacheIncrByToPipe(pipe, followingKey, -rows)
				redis_cache.CacheIncrByToPipe(pipe, fanKey, -rows)
				_, err = pipe.Exec(context.Background())
				return err
			}, followingKey, fanKey)
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &user.UnfollowUserResp{}, nil
}
