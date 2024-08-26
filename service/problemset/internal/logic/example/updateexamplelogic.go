package examplelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExampleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExampleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExampleLogic {
	return &UpdateExampleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改样例
func (l *UpdateExampleLogic) UpdateExample(in *problemset.UpdateExampleReq) (*problemset.UpdateExampleResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.UpdateExampleResp{}, nil
}
