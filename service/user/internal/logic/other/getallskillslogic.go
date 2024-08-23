package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllSkillsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSkillsLogic {
	return &GetAllSkillsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询所有技能
func (l *GetAllSkillsLogic) GetAllSkills(in *user.GetAllSkillsReq) (*user.GetAllSkillsResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetAllSkillsResp{}, nil
}
