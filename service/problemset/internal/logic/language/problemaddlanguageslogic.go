package languagelogic

import (
	"context"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemAddLanguagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProblemAddLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemAddLanguagesLogic {
	return &ProblemAddLanguagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给题目增加语言
func (l *ProblemAddLanguagesLogic) ProblemAddLanguages(in *problemset.ProblemAddLanguagesReq) (*problemset.ProblemAddLanguagesResp, error) {
	// todo: add your logic here and delete this line

	return &problemset.ProblemAddLanguagesResp{}, nil
}
