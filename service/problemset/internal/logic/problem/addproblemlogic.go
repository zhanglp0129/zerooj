package problemlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemLogic {
	return &AddProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加问题，客服权限
func (l *AddProblemLogic) AddProblem(in *problemset.AddProblemReq) (*problemset.AddProblemResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddProblemResp{}, nil
}
