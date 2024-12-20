package skilllogic

import (
	"context"
	"zerooj/service/user/models"

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

// 添加技能
func (l *AddSkillLogic) AddSkill(in *user.AddSkillReq) (*user.AddSkillResp, error) {
	// 构造数据
	s := models.Skill{
		Name: in.SkillName,
	}
	var err error
	s.ID, err = l.svcCtx.RW.GenerateId()
	if err != nil {
		return nil, err
	}
	// 插入数据
	db := l.svcCtx.DB
	err = db.Create(&s).Error
	if err != nil {
		return nil, err
	}

	return &user.AddSkillResp{
		SkillId: s.ID,
	}, nil
}
