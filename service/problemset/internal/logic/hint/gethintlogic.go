package hintlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHintLogic {
	return &GetHintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取提示
func (l *GetHintLogic) GetHint(in *problemset.GetHintReq, stream problemset.Hint_GetHintServer) error {
	// todo: add your logic here and delete this line

	return nil
}
