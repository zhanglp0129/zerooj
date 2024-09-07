package skilllogic

import (
	"context"
	"zerooj/common/constant"
	commonmodels "zerooj/common/models"
	"zerooj/service/user/models"

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
	db := l.svcCtx.DB
	// 校验技能数量
	var u models.User
	u.ID = in.UserId
	count := db.Model(&u).Association("Skills").Count()
	if int64(len(in.SkillIds))+count > 10 {
		return nil, constant.SkillQuantityExceedsLimit
	}

	// 插入用户技能
	skills := make([]models.Skill, 0, len(in.SkillIds))
	for _, skillId := range in.SkillIds {
		skills = append(skills, models.Skill{Model: commonmodels.Model{ID: skillId}})
	}
	err := db.Model(&u).Association("Skills").Append(skills)
	if err != nil {
		return nil, err
	}

	return &user.AddUserSkillResp{}, nil
}
