package registerloginlogic

import (
	"context"
	commonmodels "zerooj/common/models"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *UserRegisterLogic) UserRegister(in *user.UserRegisterReq) (*user.UserRegisterResp, error) {
	pwd, err := utils.PasswordEncrypt(in.Password)
	if err != nil {
		return nil, err
	}
	// 插入数据
	db := l.svcCtx.DB
	u := models.User{
		Username:   in.Username,
		Password:   pwd,
		Email:      in.Email,
		Permission: commonmodels.DefaultPermission,
	}
	u.ID, err = l.svcCtx.RW.GenerateId()
	if err != nil {
		return nil, err
	}
	if err = db.Create(&u).Error; err != nil {
		return nil, err
	}

	return &user.UserRegisterResp{Id: u.ID}, nil
}
