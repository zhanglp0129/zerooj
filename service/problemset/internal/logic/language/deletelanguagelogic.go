package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLanguageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLanguageLogic {
	return &DeleteLanguageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除语言，客服权限
func (l *DeleteLanguageLogic) DeleteLanguage(in *problemset.DeleteLanguageReq) (*problemset.DeleteLanguageResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.DeleteLanguageResp{}, nil
}
