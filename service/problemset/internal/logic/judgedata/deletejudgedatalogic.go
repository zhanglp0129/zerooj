package judgedatalogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJudgeDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteJudgeDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJudgeDataLogic {
	return &DeleteJudgeDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除测评数据
func (l *DeleteJudgeDataLogic) DeleteJudgeData(in *problemset.DeleteJudgeDataReq) (*problemset.DeleteJudgeDataResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteJudgeDataResp{}, nil
}
