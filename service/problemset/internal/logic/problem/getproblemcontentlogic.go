package problemlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemContentLogic {
	return &GetProblemContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取问题信息，可缓存
func (l *GetProblemContentLogic) GetProblemContent(in *problemset.GetProblemContentReq) (*problemset.GetProblemContentResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetProblemContentResp{}, nil
}
