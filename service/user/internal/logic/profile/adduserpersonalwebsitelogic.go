package profilelogic

import (
	"context"
	"zerooj/common/constant"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/models"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserPersonalWebsiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserPersonalWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserPersonalWebsiteLogic {
	return &AddUserPersonalWebsiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加个人网站，最多5个
func (l *AddUserPersonalWebsiteLogic) AddUserPersonalWebsite(in *user.AddUserPersonalWebsiteReq) (*user.AddUserPersonalWebsiteResp, error) {
	db := l.svcCtx.DB
	// 校验个人网站数量
	var count int64
	err := db.Model(&models.PersonalWebsite{}).Where("user_id = ?", in.UserId).Count(&count).Error
	if err != nil {
		return nil, err
	} else if count >= 5 {
		return nil, constant.WebsiteQuantityExceedsLimit
	}

	w := models.PersonalWebsite{
		UserID: in.UserId,
		Name:   in.Name,
		URL:    in.Url,
	}
	w.ID, err = l.svcCtx.RW.GenerateId()
	if err != nil {
		return nil, err
	}
	err = db.Create(&w).Error
	if err != nil {
		return nil, err
	}

	return &user.AddUserPersonalWebsiteResp{
		WebsiteId: w.ID,
	}, nil
}
