package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type MustDeleteTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMustDeleteTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MustDeleteTagsLogic {
	return &MustDeleteTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 强行删除标签
func (l *MustDeleteTagsLogic) MustDeleteTags(in *problemset.MustDeleteTagsReq) (*problemset.MustDeleteTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.MustDeleteTagsResp{}, nil
}
