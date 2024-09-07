package baseinfologic

import (
	"context"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/models"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBaseInfoLogic {
	return &GetBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户基本信息
func (l *GetBaseInfoLogic) GetBaseInfo(in *user.GetBaseInfoReq) (*user.GetBaseInfoResp, error) {
	return GetBaseInfo(l.svcCtx, in.Id)
}

func GetBaseInfo(svcCtx *svc.ServiceContext, userId int64) (*user.GetBaseInfoResp, error) {
	db := svcCtx.DB
	u := models.User{}
	err := db.Select("username", "email").Take(&u, userId).Error
	if err != nil {
		return nil, err
	}
	return &user.GetBaseInfoResp{
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
