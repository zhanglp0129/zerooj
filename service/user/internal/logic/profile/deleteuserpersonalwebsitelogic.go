package profilelogic

import (
	"context"
	"zerooj/common/constant"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserPersonalWebsiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserPersonalWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserPersonalWebsiteLogic {
	return &DeleteUserPersonalWebsiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除个人网站
func (l *DeleteUserPersonalWebsiteLogic) DeleteUserPersonalWebsite(in *user.DeleteUserPersonalWebsiteReq) (*user.DeleteUserPersonalWebsiteResp, error) {
	// 校验该网站是否属于该用户
	db := l.svcCtx.DB
	w := models.PersonalWebsite{}
	err := db.Select("user_id").Where("id = ?", in.WebsiteId).Take(&w).Error
	if err != nil {
		return nil, err
	} else if w.UserID != in.UserId {
		return nil, constant.InsufficientPermissionsError
	}

	err = db.Delete(&models.PersonalWebsite{}, in.WebsiteId).Error
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserPersonalWebsiteResp{}, nil
}
