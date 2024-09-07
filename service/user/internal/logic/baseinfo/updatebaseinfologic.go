package baseinfologic

import (
	"context"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseInfoLogic {
	return &UpdateBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户基本信息，字段为空表示不修改
func (l *UpdateBaseInfoLogic) UpdateBaseInfo(in *user.UpdateBaseInfoReq) (*user.UpdateBaseInfoResp, error) {
	db := l.svcCtx.DB
	var u models.User

	// 构造修改的数据模型
	encryptedPassword, err := utils.PasswordEncrypt(in.Password)
	if err != nil {
		return nil, err
	}
	u = models.User{
		Username: in.Username,
		Password: encryptedPassword,
		Email:    in.Email,
	}
	u.ID = in.Id

	// 修改数据
	err = db.Updates(&u).Error
	if err != nil {
		return nil, err
	}

	return &user.UpdateBaseInfoResp{}, nil
}
