package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLanguageByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageByIdLogic {
	return &GetLanguageByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据id获取语言
func (l *GetLanguageByIdLogic) GetLanguageById(in *problemset.GetLanguageByIdReq) (*problemset.GetLanguageByIdResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetLanguageByIdResp{}, nil
}
