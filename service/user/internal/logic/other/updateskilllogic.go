package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkillLogic {
	return &UpdateSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改技能，需要客服权限
func (l *UpdateSkillLogic) UpdateSkill(in *user.UpdateSkillReq) (*user.UpdateSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateSkillResp{}, nil
}
