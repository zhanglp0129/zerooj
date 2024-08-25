package submitlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitCodeLogic {
	return &SubmitCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户提交代码
func (l *SubmitCodeLogic) SubmitCode(in *problemset.SubmitCodeReq) (*problemset.SubmitCodeResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.SubmitCodeResp{}, nil
}
