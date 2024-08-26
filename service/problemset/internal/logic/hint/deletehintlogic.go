package hintlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHintLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHintLogic {
	return &DeleteHintLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除提示
func (l *DeleteHintLogic) DeleteHint(in *problemset.DeleteHintReq) (*problemset.DeleteHintResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteHintResp{}, nil
}
