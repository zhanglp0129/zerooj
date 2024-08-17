package baseinfologic

import (
	"context"
	"fmt"
	"zerooj/common"
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
	m := user.GetBaseInfoResp{}
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("/cache/user/get_base_info/%d", in.Id)
	// 带着缓存查询
	err := common.WithCache(rdb, key, &m, func() (*user.GetBaseInfoResp, error) {
		db := l.svcCtx.DB
		u := models.User{}
		err := db.Take(&u, in.Id).Error
		if err != nil {
			return nil, err
		}
		return &user.GetBaseInfoResp{
			Username:   u.Username,
			Email:      u.Email,
			Permission: uint32(u.Permission),
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return &m, nil
}
