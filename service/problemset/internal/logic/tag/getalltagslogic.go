package taglogic

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTagsLogic {
	return &GetAllTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有标签，可缓存
func (l *GetAllTagsLogic) GetAllTags(in *problemset.GetAllTagsReq) (*problemset.GetAllTagsResp, error) {
	rdb := l.svcCtx.RDB
	key := "cache:all_tags"
	var model []*problemset.TagInfo
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*[]*problemset.TagInfo, error) {
		var tags []models.Tag
		db := l.svcCtx.DB
		err := db.Find(&tags).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		result := make([]*problemset.TagInfo, len(tags))
		for i, tag := range tags {
			result[i].TagId = tag.ID
			result[i].Name = tag.Name
		}

		return &result, nil
	})
	if err != nil {
		return nil, err
	}

	return &problemset.GetAllTagsResp{
		Tags: model,
	}, nil
}

// DeleteAllTagsCache 删除所有标签的缓存
func DeleteAllTagsCache(rdb redis.UniversalClient) error {
	key := "cache:all_tags"
	return redis_cache.DeleteCache(rdb, key)
}
