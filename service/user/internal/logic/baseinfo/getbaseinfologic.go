package baseinfologic

import (
	"context"
	"fmt"
	"github.com/zhanglp0129/redis_cache"
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

// 获取用户基本信息，不包括密码，并缓存
func (l *GetBaseInfoLogic) GetBaseInfo(in *user.GetBaseInfoReq) (*user.GetBaseInfoResp, error) {
	return GetBaseInfo(l.svcCtx, in.Id)
}

// GetBaseInfo 获取用户基础信息。方便复用代码
func GetBaseInfo(svcCtx *svc.ServiceContext, userId int64) (*user.GetBaseInfoResp, error) {
	// 带着缓存查询
	key := fmt.Sprintf("cache:user_base_info:%d", userId)
	rdb := svcCtx.RDB
	model := user.GetBaseInfoResp{}
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*user.GetBaseInfoResp, error) {
		db := svcCtx.DB
		u := models.User{}
		err := db.Take(&u, userId).Error
		if err != nil {
			return nil, err
		}
		return &user.GetBaseInfoResp{
			Username: u.Username,
			Email:    u.Email,
		}, nil
	})

	return &model, err
}
