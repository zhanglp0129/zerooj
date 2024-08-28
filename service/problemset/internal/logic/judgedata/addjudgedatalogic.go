package judgedatalogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJudgeDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJudgeDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJudgeDataLogic {
	return &AddJudgeDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加测评数据
func (l *AddJudgeDataLogic) AddJudgeData(stream problemset.JudgeData_AddJudgeDataServer) error {
	// todo: add your logic here and delete this line

	return nil
}
