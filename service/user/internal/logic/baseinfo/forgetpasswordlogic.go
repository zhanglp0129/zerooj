package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 忘记密码，需重新登录
func (l *ForgetPasswordLogic) ForgetPassword(in *user.ForgetPasswordReq) (*user.ForgetPasswordResp, error) {
	// todo: add your logic here and delete this line

	return &user.ForgetPasswordResp{}, nil
}
