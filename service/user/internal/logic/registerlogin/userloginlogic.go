package registerloginlogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"zerooj/common/constant"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *UserLoginLogic) UserLogin(in *user.UserLoginReq) (*user.UserLoginResp, error) {
	pwd, err := utils.PasswordEncrypt(in.Password)
	if err != nil {
		return nil, err
	}

	db := l.svcCtx.DB
	u := &models.User{}
	if in.Username == "" {
		// 邮箱登录
		err = db.Where("email = ? and password = ?", in.Email, pwd).Take(u).Error
	} else {
		// 用户名登录
		err = db.Where("username = ? and password = ?", in.Username, pwd).Take(u).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有记录
		return nil, constant.UsernameEmailPasswordError{}
	} else if err != nil {
		return nil, err
	}

	return &user.UserLoginResp{
		Id:         u.ID,
		Permission: uint32(u.Permission),
	}, nil
}
