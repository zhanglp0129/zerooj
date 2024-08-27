package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemLanguagesLogic {
	return &GetProblemLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取题目的所有语言
func (l *GetProblemLanguagesLogic) GetProblemLanguages(in *problemset.GetProblemLanguagesReq) (*problemset.GetProblemLanguagesResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.GetProblemLanguagesResp{}, nil
}
