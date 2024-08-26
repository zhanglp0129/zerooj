package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagsLogic {
	return &DeleteTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除标签
func (l *DeleteTagsLogic) DeleteTags(in *problemset.DeleteTagsReq) (*problemset.DeleteTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteTagsResp{}, nil
}
