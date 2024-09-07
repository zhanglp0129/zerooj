package baseinfologic

import (
	"context"
	"fmt"
	"github.com/zhanglp0129/redis_cache"
	"zerooj/service/user/models"

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

// 获取用户权限，并缓存
func (l *GetPermissionLogic) GetPermission(in *user.GetPermissionReq) (*user.GetPermissionResp, error) {
	return GetPermission(l.svcCtx, in.Id)
}

// GetPermission 获取用户权限。方便复用代码
func GetPermission(svcCtx *svc.ServiceContext, userId int64) (*user.GetPermissionResp, error) {
	// 带着缓存查询
	key := fmt.Sprintf("cache:user_permission:%d", userId)
	rdb := svcCtx.RDB
	model := user.GetPermissionResp{}
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*user.GetPermissionResp, error) {
		db := svcCtx.DB
		u := models.User{}
		err := db.Select("permission").Take(&u, userId).Error
		if err != nil {
			return nil, err
		}
		return &user.GetPermissionResp{
			Permission: uint32(u.Permission),
		}, nil
	})

	return &model, err
}
