package otherlogic

import (
	"context"
	"zerooj/common/constant"
	common_models "zerooj/common/models"
	baseinfologic "zerooj/service/user/internal/logic/baseinfo"
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

// 修改技能，需要客服权限
func (l *UpdateSkillLogic) UpdateSkill(in *user.UpdateSkillReq) (*user.UpdateSkillResp, error) {
	// 鉴权
	baseInfo, err := baseinfologic.GetBaseInfo(l.svcCtx, in.OperatorId)
	if err != nil {
		return nil, err
	}
	if !common_models.Permission(baseInfo.Permission).CanSupport() {
		return nil, constant.InsufficientPermissionsError{}
	}

	db := l.svcCtx.DB
	err = db.Model(&models.Skill{}).Where("id = ?", in.SkillId).Update("name", in.SkillName).Error
	if err != nil {
		return nil, err
	}
	return &user.UpdateSkillResp{}, nil
}
