package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemDeleteTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProblemDeleteTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemDeleteTagsLogic {
	return &ProblemDeleteTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给题目删除标签，客服权限
func (l *ProblemDeleteTagsLogic) ProblemDeleteTags(in *problemset.ProblemDeleteTagsReq) (*problemset.ProblemDeleteTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.ProblemDeleteTagsResp{}, nil
}
