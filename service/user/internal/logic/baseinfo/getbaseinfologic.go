package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
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

// 获取用户基本信息，不包括密码
func (l *GetBaseInfoLogic) GetBaseInfo(in *user.GetBaseInfoReq) (*user.GetBaseInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetBaseInfoResp{}, nil
}
