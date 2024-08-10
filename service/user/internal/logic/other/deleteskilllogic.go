package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkillLogic {
	return &DeleteSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSkillLogic) DeleteSkill(in *user.DeleteSkillReq) (*user.DeleteSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteSkillResp{}, nil
}
