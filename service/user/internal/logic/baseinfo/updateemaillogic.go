package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户邮箱，需重新登录，有30天冷却期
func (l *UpdateEmailLogic) UpdateEmail(in *user.UpdateEmailReq) (*user.UpdateEmailResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateEmailResp{}, nil
}
