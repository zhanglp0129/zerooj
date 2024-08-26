package hintlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHintLogic {
	return &AddHintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加提示
func (l *AddHintLogic) AddHint(in *problemset.AddHintReq) (*problemset.AddHintResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddHintResp{}, nil
}
