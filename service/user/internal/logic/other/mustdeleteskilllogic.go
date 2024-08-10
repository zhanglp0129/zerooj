package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MustDeleteSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMustDeleteSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MustDeleteSkillLogic {
	return &MustDeleteSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 强行删除技能，必须要管理员权限
func (l *MustDeleteSkillLogic) MustDeleteSkill(in *user.MustDeleteSkillReq) (*user.MustDeleteSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.MustDeleteSkillResp{}, nil
}
