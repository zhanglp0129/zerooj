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

// 添加测评数据，客服权限
func (l *AddJudgeDataLogic) AddJudgeData(in *problemset.AddJudgeDataReq) (*problemset.AddJudgeDataResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddJudgeDataResp{}, nil
}
