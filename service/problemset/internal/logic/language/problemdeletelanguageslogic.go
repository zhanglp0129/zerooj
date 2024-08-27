package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemDeleteLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProblemDeleteLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemDeleteLanguagesLogic {
	return &ProblemDeleteLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 题目删除语言
func (l *ProblemDeleteLanguagesLogic) ProblemDeleteLanguages(in *problemset.ProblemDeleteLanguagesReq) (*problemset.ProblemDeleteLanguagesResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.ProblemDeleteLanguagesResp{}, nil
}
