package judgedatalogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgeDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeDataLogic {
	return &GetJudgeDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取题目的测评数据，返回minio对象名称，可缓存
func (l *GetJudgeDataLogic) GetJudgeData(in *problemset.GetJudgeDataReq, stream problemset.JudgeData_GetJudgeDataServer) error {
	// todo: add your logic here and delete this line

	return nil
}
