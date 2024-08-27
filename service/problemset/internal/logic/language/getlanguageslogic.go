package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有语言
func (l *GetLanguagesLogic) GetLanguages(in *problemset.GetLanguagesReq) (*problemset.GetLanguagesResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetLanguagesResp{}, nil
}
