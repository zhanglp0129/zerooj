package submitlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmitByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmitByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmitByIdLogic {
	return &GetSubmitByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取通过id提交记录
func (l *GetSubmitByIdLogic) GetSubmitById(in *problemset.GetSubmitByIdReq) (*problemset.GetSubmitByIdResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetSubmitByIdResp{}, nil
}
