package judgelogic

import (
	"context"

	"zerooj/service/judge/internal/svc"
	"zerooj/service/judge/pb/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeWithDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeWithDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeWithDataLogic {
	return &JudgeWithDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 测评，并提供测评数据
func (l *JudgeWithDataLogic) JudgeWithData(stream judge.Judge_JudgeWithDataServer) error {
	// todo: add your logic here and delete this line

	return nil
}
