package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagsLogic {
	return &AddTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加标签，客服权限
func (l *AddTagsLogic) AddTags(in *problemset.AddTagsReq) (*problemset.AddTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddTagsResp{}, nil
}
