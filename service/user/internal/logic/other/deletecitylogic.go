package otherlogic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCityLogic {
	return &DeleteCityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCityLogic) DeleteCity(in *user.DeleteCityReq) (*user.DeleteCityResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteCityResp{}, nil
}
