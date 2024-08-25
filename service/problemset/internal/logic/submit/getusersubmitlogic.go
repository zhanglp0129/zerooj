package submitlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSubmitLogic {
	return &GetUserSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取用户提交
func (l *GetUserSubmitLogic) GetUserSubmit(in *problemset.GetUserSubmitReq) (*problemset.GetUserSubmitResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetUserSubmitResp{}, nil
}
