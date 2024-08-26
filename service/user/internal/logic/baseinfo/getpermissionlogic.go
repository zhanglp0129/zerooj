package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户权限
func (l *GetPermissionLogic) GetPermission(in *user.GetPermissionReq) (*user.GetPermissionResp, error) {
	// 获取用户基本信息
	baseInfo, err := GetBaseInfo(l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetPermissionResp{
		Permission: baseInfo.Permission,
	}, nil
}
