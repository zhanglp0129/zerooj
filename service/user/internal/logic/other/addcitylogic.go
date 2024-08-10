package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCityLogic {
	return &AddCityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加和删除城市，只有客服能操作
func (l *AddCityLogic) AddCity(in *user.AddCityReq) (*user.AddCityResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddCityResp{}, nil
}
