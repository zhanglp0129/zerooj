package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsernameLogic {
	return &UpdateUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户名，需重新登录，有30天冷却期
func (l *UpdateUsernameLogic) UpdateUsername(in *user.UpdateUsernameReq) (*user.UpdateUsernameResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateUsernameResp{}, nil
}
