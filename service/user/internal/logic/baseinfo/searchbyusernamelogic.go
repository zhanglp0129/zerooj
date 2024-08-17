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

type SearchByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByUsernameLogic {
	return &SearchByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户名搜索用户，并缓存
func (l *SearchByUsernameLogic) SearchByUsername(in *user.SearchByUsernameReq) (*user.SearchByUsernameResp, error) {
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf(" /cache/user/search_by_username/%s", in.Username)
	model := user.SearchByUsernameResp{}

	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*user.SearchByUsernameResp, error) {
		db := l.svcCtx.DB
		u := models.User{}
		err := db.Select("id").Where("username = ?", in.Username).Take(&u).Error
		if err != nil {
			return nil, err
		}

		return &user.SearchByUsernameResp{
			Id: u.ID,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return &model, nil
}
