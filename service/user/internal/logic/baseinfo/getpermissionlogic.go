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

// 获取用户权限。鉴权时优先使用该接口
func (l *GetPermissionLogic) GetPermission(in *user.GetPermissionReq) (*user.GetPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetPermissionResp{}, nil
}
