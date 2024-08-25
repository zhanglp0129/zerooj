package submitlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemSubmitCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemSubmitCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemSubmitCountLogic {
	return &GetProblemSubmitCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取题目的提交数
func (l *GetProblemSubmitCountLogic) GetProblemSubmitCount(in *problemset.GetProblemSubmitCountReq) (*problemset.GetProblemSubmitCountResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetProblemSubmitCountResp{}, nil
}
