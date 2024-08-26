package otherlogic

import (
	"context"
	"gorm.io/gorm"
	"zerooj/service/user/models"

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

// 强行删除技能
func (l *MustDeleteSkillLogic) MustDeleteSkill(in *user.MustDeleteSkillReq) (*user.MustDeleteSkillResp, error) {
	var count int64
	db := l.svcCtx.DB
	err := db.Transaction(func(tx *gorm.DB) error {
		// 先删除所有拥有该技能的用户
		var skill models.Skill
		skill.ID = in.SkillId
		tx2 := tx.Model(&skill).Association("Users")
		count = tx2.Count()
		err := tx2.Clear()
		if err != nil {
			return err
		}

		// 删除该技能
		err = tx.Delete(&models.Skill{}, in.SkillId).Error
		return err
	})
	if err != nil {
		return nil, err
	}

	return &user.MustDeleteSkillResp{
		UserSkillCount: count,
	}, nil
}
