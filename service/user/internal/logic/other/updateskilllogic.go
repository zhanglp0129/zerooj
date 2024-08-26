package otherlogic

import (
	"context"
	"zerooj/service/user/models"

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

// 修改技能
func (l *UpdateSkillLogic) UpdateSkill(in *user.UpdateSkillReq) (*user.UpdateSkillResp, error) {
	db := l.svcCtx.DB
	err := db.Model(&models.Skill{}).Where("id = ?", in.SkillId).Update("name", in.SkillName).Error
	if err != nil {
		return nil, err
	}
	return &user.UpdateSkillResp{}, nil
}
