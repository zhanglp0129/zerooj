package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type MustDeleteCityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMustDeleteCityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MustDeleteCityLogic {
	return &MustDeleteCityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 强行删除城市，必须要管理员权限
func (l *MustDeleteCityLogic) MustDeleteCity(in *user.MustDeleteCityReq) (*user.MustDeleteCityResp, error) {
	// todo: add your logic here and delete this line

	return &user.MustDeleteCityResp{}, nil
}
