package taglogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type MustDeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMustDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MustDeleteTagLogic {
	return &MustDeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 强行删除标签
func (l *MustDeleteTagLogic) MustDeleteTag(in *problemset.MustDeleteTagReq) (*problemset.MustDeleteTagResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.MustDeleteTagResp{}, nil
}
