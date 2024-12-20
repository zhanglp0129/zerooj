package baseinfologic

import (
	"context"
	"fmt"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"zerooj/service/user/models"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户权限，需要删除缓存
func (l *UpdatePermissionLogic) UpdatePermission(in *user.UpdatePermissionReq) (*user.UpdatePermissionResp, error) {
	db := l.svcCtx.DB
	err := db.Transaction(func(tx *gorm.DB) error {
		// 修改数据
		err := tx.Model(&models.User{}).Where("id = ?", in.Id).Update("permission", in.Permission).Error
		if err != nil {
			return err
		}

		// 删除缓存
		rdb := l.svcCtx.RDB
		getBaseInfoKey := fmt.Sprintf("cache:user_permission:%d", in.Id)
		return redis_cache.DeleteCache(rdb, getBaseInfoKey)
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdatePermissionResp{}, nil
}
