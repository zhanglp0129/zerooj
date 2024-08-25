package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLanguageLogic {
	return &UpdateLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新语言，客服权限
func (l *UpdateLanguageLogic) UpdateLanguage(in *problemset.UpdateLanguageReq) (*problemset.UpdateLanguageResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.UpdateLanguageResp{}, nil
}
