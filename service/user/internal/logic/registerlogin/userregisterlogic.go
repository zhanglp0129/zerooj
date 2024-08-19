package registerloginlogic

import (
	"context"
	"fmt"
	"zerooj/common"
	"zerooj/common/constant"
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
	// 先校验邮箱验证码
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("/mail_check_code/user/user_register/%s", in.Email)
	trueCheckCode, err := rdb.Get(context.Background(), key).Result()
	if err != nil || in.EmailCheckCode != trueCheckCode {
		return nil, constant.InputDataError{Thing: "邮箱验证码"}
	}

	pwd, err := utils.PasswordEncrypt(in.Password)
	if err != nil {
		return nil, err
	}
	// 插入数据
	db := l.svcCtx.DB
	u := &models.User{
		Username:   in.Username,
		Password:   pwd,
		Email:      in.Email,
		Permission: common.DefaultPermission,
	}
	if err = db.Create(u).Error; err != nil {
		return nil, err
	}

	common.FinishMailCheckCode(rdb, key)

	return &user.UserRegisterResp{Id: u.ID}, nil
}
