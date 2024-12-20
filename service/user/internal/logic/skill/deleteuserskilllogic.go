package skilllogic

import (
	"context"
	commonmodels "zerooj/common/models"
	"zerooj/service/user/models"

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
	db := l.svcCtx.DB
	var u models.User
	u.ID = in.UserId
	skills := make([]models.Skill, 0, len(in.SkillIds))
	for _, skillId := range in.SkillIds {
		skills = append(skills, models.Skill{Model: commonmodels.Model{ID: skillId}})
	}
	err := db.Model(&u).Association("Skills").Delete(&skills)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserSkillResp{}, nil
}
