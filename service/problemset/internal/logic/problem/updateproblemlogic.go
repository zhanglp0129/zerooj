package problemlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新问题
func (l *UpdateProblemLogic) UpdateProblem(in *problemset.UpdateProblemReq) (*problemset.UpdateProblemResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.UpdateProblemResp{}, nil
}
