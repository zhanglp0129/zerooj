package profilelogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSkillLogic {
	return &AddUserSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户技能，最多10个
func (l *AddUserSkillLogic) AddUserSkill(in *user.AddUserSkillReq) (*user.AddUserSkillResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddUserSkillResp{}, nil
}
