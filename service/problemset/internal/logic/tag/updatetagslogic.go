package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagsLogic {
	return &UpdateTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新标签
func (l *UpdateTagsLogic) UpdateTags(in *problemset.UpdateTagsReq) (*problemset.UpdateTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.UpdateTagsResp{}, nil
}
