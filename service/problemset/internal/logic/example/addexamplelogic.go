package examplelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddExampleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddExampleLogic {
	return &AddExampleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加样例
func (l *AddExampleLogic) AddExample(in *problemset.AddExampleReq) (*problemset.AddExampleResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddExampleResp{}, nil
}
