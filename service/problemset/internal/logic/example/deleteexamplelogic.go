package examplelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExampleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExampleLogic {
	return &DeleteExampleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除样例
func (l *DeleteExampleLogic) DeleteExample(in *problemset.DeleteExampleReq) (*problemset.DeleteExampleResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteExampleResp{}, nil
}
