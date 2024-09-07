package profilelogic

import (
	"context"
	commonmodels "zerooj/common/models"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户简介，不包括个人网站和用户技能
func (l *UpdateUserProfileLogic) UpdateUserProfile(in *user.UpdateUserProfileReq) (*user.UpdateUserProfileResp, error) {
	db := l.svcCtx.DB
	p := models.UserProfile{
		UserID:               in.UserId,
		Nickname:             in.Nickname,
		AvatarURL:            in.AvatarURL,
		Gender:               commonmodels.Gender(in.Gender),
		Birthday:             in.Birthday.AsTime(),
		PersonalIntroduction: in.PersonalIntroduction,
	}
	err := db.Select("*").Omit("ID", "CreatedAt", "UserID").Where("user_id = ?", in.UserId).Updates(&p).Error
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserProfileResp{}, nil
}
