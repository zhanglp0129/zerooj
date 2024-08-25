package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLanguageLogic {
	return &AddLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加语言，客服权限
func (l *AddLanguageLogic) AddLanguage(in *problemset.AddLanguageReq) (*problemset.AddLanguageResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.AddLanguageResp{}, nil
}
