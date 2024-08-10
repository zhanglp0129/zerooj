package profilelogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户简介，不包括个人网站和用户技能
func (l *UpdateUserProfileLogic) UpdateUserProfile(in *user.UpdateUserProfileReq) (*user.UpdateUserProfileResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateUserProfileResp{}, nil
}
