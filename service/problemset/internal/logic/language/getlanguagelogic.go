package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageLogic {
	return &GetLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有语言
func (l *GetLanguageLogic) GetLanguage(in *problemset.GetLanguageReq) (*problemset.GetLanguageResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetLanguageResp{}, nil
}
