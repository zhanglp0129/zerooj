package judgedatalogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJudgeDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateJudgeDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJudgeDataLogic {
	return &UpdateJudgeDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改测评数据
func (l *UpdateJudgeDataLogic) UpdateJudgeData(stream problemset.JudgeData_UpdateJudgeDataServer) error {
	// todo: add your logic here and delete this line

	return nil
}
