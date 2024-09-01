package judgelogic

import (
	"context"

	"zerooj/service/judge/internal/svc"
	"zerooj/service/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 测评
func (l *JudgeLogic) Judge(stream judge.Judge_JudgeServer) error {
	// todo: add your logic here and delete this line

	return nil
}
