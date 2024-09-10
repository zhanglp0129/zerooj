package taglogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"gorm.io/gorm"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemTagsLogic {
	return &GetProblemTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取一个题目的所有标签，可缓存
func (l *GetProblemTagsLogic) GetProblemTags(in *problemset.GetProblemTagsReq) (*problemset.GetProblemTagsResp, error) {
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("cache:problem_tags:%d", in.ProblemId)
	var model problemset.GetProblemTagsResp
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*problemset.GetProblemTagsResp, error) {
		db := l.svcCtx.DB
		var tags []models.Tag
		err := db.Model(&models.Problem{}).Where("id = ?", in.ProblemId).
			Association("Tags").Find(&tags)
		if errors.Is(err, gorm.ErrRecordNotFound) || len(tags) == 0 {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		var result problemset.GetProblemTagsResp
		result.Tags = make([]*problemset.TagInfo, len(tags))
		for i := range tags {
			result.Tags[i] = &problemset.TagInfo{
				TagId: tags[i].ID,
				Name:  tags[i].Name,
			}

		}

		return &result, nil
	})
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// DeleteProblemTagsCache 删除题目标签缓存
func DeleteProblemTagsCache(rdb redis.UniversalClient, problemId int64) error {
	key := fmt.Sprintf("cache:problem_tags:%d", problemId)
	return redis_cache.DeleteCache(rdb, key)
}
