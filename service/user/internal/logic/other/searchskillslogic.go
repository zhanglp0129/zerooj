package otherlogic

import (
	"context"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSkillsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSkillsLogic {
	return &SearchSkillsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索技能
func (l *SearchSkillsLogic) SearchSkills(in *user.SearchSkillsReq) (*user.SearchSkillsResp, error) {
	// 准备查询
	db := l.svcCtx.DB
	ss := make([]models.Skill, 0)
	offset := (in.PageNum - 1) * in.PageSize
	tx := db.Model(&models.Skill{}).Where("name like ?", "%"+in.Name+"%")
	// 查询总数
	var count int64
	err := tx.Count(&count).Error
	if err != nil {
		return nil, err
	}
	// 查询结果
	tx = tx.Offset(int(offset)).Limit(int(in.PageSize)).Order("name")
	err = tx.Find(&ss).Error
	if err != nil {
		return nil, err
	}

	skills := make([]*user.Skill, 0, len(ss))
	for _, s := range ss {
		skills = append(skills, &user.Skill{
			SkillId: s.ID,
			Name:    s.Name,
		})
	}

	return &user.SearchSkillsResp{
		Skills: skills,
		Count:  count,
	}, nil
}
