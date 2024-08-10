package profilelogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserSkillLogic {
	return &DeleteUserSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户技能
func (l *DeleteUserSkillLogic) DeleteUserSkill(in *user.DeleteUserSkillReq) (*user.DeleteUserSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteUserSkillResp{}, nil
}
