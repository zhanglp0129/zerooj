package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSkillLogic {
	return &AddSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加和删除技能，只有客服能操作
func (l *AddSkillLogic) AddSkill(in *user.AddSkillReq) (*user.AddSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddSkillResp{}, nil
}
