package baseinfologic

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

type ForgetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 忘记密码
func (l *ForgetPasswordLogic) ForgetPassword(in *user.ForgetPasswordReq) (*user.ForgetPasswordResp, error) {
	// 查询邮箱并验证邮箱验证码
	db := l.svcCtx.DB
	u := models.User{}
	err := db.Select("email", "password").Take(&u, in.Id).Error
	if err != nil {
		return nil, err
	}
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("/mail_check_code/user/forget_password/%s", u.Email)
	trueCheckCode, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	if trueCheckCode != in.EmailCheckCode {
		return nil, constant.InputDataError{Thing: "邮箱验证码"}
	}

	// 检查新旧密码是否相同
	newPasswordEncrypt, err := utils.PasswordEncrypt(in.NewPassword)
	if err != nil {
		return nil, err
	}
	if newPasswordEncrypt == u.Password {
		return nil, constant.NewEqualsOldError{Thing: "密码"}
	}

	// 修改密码
	err = db.Model(&u).Where("id = ?", in.Id).Update("password", newPasswordEncrypt).Error
	if err != nil {
		return nil, err
	}

	go common.FinishMailCheckCode(rdb, key)

	return &user.ForgetPasswordResp{}, nil
}
