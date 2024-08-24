package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTagsLogic {
	return &GetAllTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有标签，可缓存
func (l *GetAllTagsLogic) GetAllTags(in *problemset.GetAllTagsReq) (*problemset.GetAllTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetAllTagsResp{}, nil
}
