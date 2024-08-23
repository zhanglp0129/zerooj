package followlogic

import (
	"context"
	"fmt"
	"github.com/zhanglp0129/redis_cache"
	"google.golang.org/protobuf/types/known/timestamppb"
	baseinfologic "zerooj/service/user/internal/logic/baseinfo"
	profilelogic "zerooj/service/user/internal/logic/profile"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingsLogic {
	return &GetFollowingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有关注，分页查询
func (l *GetFollowingsLogic) GetFollowings(in *user.GetFollowingsReq) (*user.GetFollowingsResp, error) {
	db := l.svcCtx.DB
	rdb := l.svcCtx.RDB
	var count int64
	key := fmt.Sprintf("/cache/user/get_followings/count/%d", in.UserId)
	// 获取总数
	_, err := redis_cache.QueryWithCache(rdb, key, &count, func() (*int64, error) {
		var total int64
		err := db.Model(&models.Follow{}).Where("follower_id = ?", in.UserId).Count(&total).Error
		if err != nil {
			return nil, err
		}
		return &total, nil
	})
	if err != nil {
		return nil, err
	}

	// 分页查询
	var offset = (in.PageNum - 1) * in.PageSize
	follows := make([]models.Follow, 0)
	err = db.Where("follower_id = ?", in.UserId).Offset(int(offset)).Limit(int(in.PageSize)).Find(&follows).Error
	if err != nil {
		return nil, err
	}
	followings := make([]*user.FollowUserInfo, 0, len(follows))
	for _, follow := range follows {
		baseInfo, err := baseinfologic.GetBaseInfo(l.svcCtx, follow.FollowedID)
		if err != nil {
			return nil, err
		}
		profile, err := profilelogic.GetUserProfile(l.svcCtx, follow.FollowedID)
		if err != nil {
			return nil, err
		}
		followings = append(followings, &user.FollowUserInfo{
			Id:                   follow.FollowedID,
			Username:             baseInfo.Username,
			Nickname:             profile.Nickname,
			PersonalIntroduction: profile.PersonalIntroduction,
			AvatarURL:            profile.AvatarURL,
			Websites:             profile.Websites,
			Skills:               profile.Skills,
			FollowAt:             timestamppb.New(follow.FollowAt),
		})
	}

	return &user.GetFollowingsResp{
		Count:      count,
		Followings: followings,
	}, nil
}
