package hintlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemHintsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemHintsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemHintsLogic {
	return &GetProblemHintsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取题目所有提示
func (l *GetProblemHintsLogic) GetProblemHints(in *problemset.GetProblemHintsReq) (*problemset.GetProblemHintsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetProblemHintsResp{}, nil
}
