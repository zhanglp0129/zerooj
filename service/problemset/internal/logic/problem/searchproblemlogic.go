package problemlogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemLogic {
	return &SearchProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页搜索题目
func (l *SearchProblemLogic) SearchProblem(in *problemset.SearchProblemReq) (*problemset.SearchProblemResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.SearchProblemResp{}, nil
}
