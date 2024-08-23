package followlogic

import (
	"context"
	"errors"
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

	err := db.Create(&follow).Error
	if err != nil {
		return nil, err
	}

	return &user.FollowUserResp{}, nil
}
