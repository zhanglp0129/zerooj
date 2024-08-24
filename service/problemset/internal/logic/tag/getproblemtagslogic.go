package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemTagsLogic {
	return &GetProblemTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取一个题目的所有标签，可缓存
func (l *GetProblemTagsLogic) GetProblemTags(in *problemset.GetProblemTagsReq) (*problemset.GetProblemTagsResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetProblemTagsResp{}, nil
}
