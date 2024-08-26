package problemlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProblemLogic {
	return &DeleteProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除问题
func (l *DeleteProblemLogic) DeleteProblem(in *problemset.DeleteProblemReq) (*problemset.DeleteProblemResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteProblemResp{}, nil
}
