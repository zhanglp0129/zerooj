package examplelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemExamplesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemExamplesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemExamplesLogic {
	return &GetProblemExamplesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取题目的所有样例
func (l *GetProblemExamplesLogic) GetProblemExamples(in *problemset.GetProblemExamplesReq, stream problemset.Example_GetProblemExamplesServer) error {
	// todo: add your logic here and delete this line

	return nil
}
