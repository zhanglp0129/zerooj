package profilelogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserPersonalWebsiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserPersonalWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserPersonalWebsiteLogic {
	return &AddUserPersonalWebsiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加个人网站，最多5个
func (l *AddUserPersonalWebsiteLogic) AddUserPersonalWebsite(in *user.AddUserPersonalWebsiteReq) (*user.AddUserPersonalWebsiteResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddUserPersonalWebsiteResp{}, nil
}
