package followlogic

import (
	"context"
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

	err := db.Delete(&follow).Error
	if err != nil {
		return nil, err
	}

	return &user.UnfollowUserResp{}, nil
}
