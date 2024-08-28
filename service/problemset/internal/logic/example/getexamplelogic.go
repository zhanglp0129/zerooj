package examplelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExampleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExampleLogic {
	return &GetExampleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取样例
func (l *GetExampleLogic) GetExample(in *problemset.GetExampleReq, stream problemset.Example_GetExampleServer) error {
	// todo: add your logic here and delete this line

	return nil
}
