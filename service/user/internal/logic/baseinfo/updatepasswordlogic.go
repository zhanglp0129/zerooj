package baseinfologic

import (
	"context"
	"zerooj/common/constant"
	"zerooj/service/user/models"
	"zerooj/utils"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改密码
func (l *UpdatePasswordLogic) UpdatePassword(in *user.UpdatePasswordReq) (*user.UpdatePasswordResp, error) {
	// 先校验原密码
	db := l.svcCtx.DB
	u := models.User{}
	err := db.Select("password").Take(&u, in.Id).Error
	if err != nil {
		return nil, err
	}
	oldPasswordEncrypt, err := utils.PasswordEncrypt(in.OldPassword)
	if err != nil {
		return nil, err
	} else if oldPasswordEncrypt != u.Password {
		return nil, constant.InputDataError{Thing: "原密码"}
	}

	// 再校验新旧密码是否相同
	if in.OldPassword == in.NewPassword {
		return nil, constant.NewEqualsOldError{Thing: "密码"}
	}

	//  修改密码
	newPasswordEncrypt, err := utils.PasswordEncrypt(in.NewPassword)
	if err != nil {
		return nil, err
	}
	err = db.Model(&u).Where("id = ?", in.Id).Update("password", newPasswordEncrypt).Error
	if err != nil {
		return nil, err
	}

	return &user.UpdatePasswordResp{}, nil
}
