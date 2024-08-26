package otherlogic

import (
	"context"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkillLogic {
	return &DeleteSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSkillLogic) DeleteSkill(in *user.DeleteSkillReq) (*user.DeleteSkillResp, error) {
	// 删除
	db := l.svcCtx.DB
	err := db.Delete(&models.Skill{}, in.SkillId).Error
	if err != nil {
		return nil, err
	}

	return &user.DeleteSkillResp{}, nil
}
