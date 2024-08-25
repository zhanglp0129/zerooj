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
	// 鉴权
	baseInfo, err := baseinfologic.GetBaseInfo(l.svcCtx, in.OperatorId)
	if err != nil {
		return nil, err
	}
	if !common_models.Permission(baseInfo.Permission).CanSupport() {
		return nil, constant.InsufficientPermissionsError
	}

	// 删除
	db := l.svcCtx.DB
	err = db.Delete(&models.Skill{}, in.SkillId).Error
	if err != nil {
		return nil, err
	}

	return &user.DeleteSkillResp{}, nil
}
