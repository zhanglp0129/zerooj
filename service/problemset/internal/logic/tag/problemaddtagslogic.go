package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemAddTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProblemAddTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemAddTagsLogic {
	return &ProblemAddTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给题目添加标签，最多5个
func (l *ProblemAddTagsLogic) ProblemAddTags(in *problemset.ProblemAddTagsReq) (*problemset.ProblemAddTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.ProblemAddTagsResp{}, nil
}
