package profilelogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserPersonalWebsiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserPersonalWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserPersonalWebsiteLogic {
	return &DeleteUserPersonalWebsiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除个人网站
func (l *DeleteUserPersonalWebsiteLogic) DeleteUserPersonalWebsite(in *user.DeleteUserPersonalWebsiteReq) (*user.DeleteUserPersonalWebsiteResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserPersonalWebsiteResp{}, nil
}
