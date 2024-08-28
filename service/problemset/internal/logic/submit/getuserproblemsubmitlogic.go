package submitlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProblemSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProblemSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProblemSubmitLogic {
	return &GetUserProblemSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户某一道题的全部提交
func (l *GetUserProblemSubmitLogic) GetUserProblemSubmit(in *problemset.GetUserProblemSubmitReq, stream problemset.Submit_GetUserProblemSubmitServer) error {
	// todo: add your logic here and delete this line

	return nil
}
