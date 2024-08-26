package hintlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHintLogic {
	return &UpdateHintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改提示
func (l *UpdateHintLogic) UpdateHint(in *problemset.UpdateHintReq) (*problemset.UpdateHintResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.UpdateHintResp{}, nil
}
