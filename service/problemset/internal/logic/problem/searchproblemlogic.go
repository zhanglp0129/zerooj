package problemlogic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProblemLogic {
	return &SearchProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页搜索题目
func (l *SearchProblemLogic) SearchProblem(in *problemset.SearchProblemReq) (*problemset.SearchProblemResp, error) {
	// 获取查询条件的tx
	db := l.svcCtx.DB
	tx := db.Model(&models.Problem{})
	if in.ProblemTitle != "" {
		tx = tx.Where("title like ?", fmt.Sprintf("%%%s%%", in.ProblemTitle))
	}
	if in.Difficulty != 0 {
		tx = tx.Where("difficulty = ?", in.Difficulty)
	}
	if len(in.TagIds) != 0 {
		tx = tx.Preload("Tags", "tags.id in (?)", in.TagIds).Group("problems.id")
	}

	var count int64
	queryCount := func() (*int64, error) {
		var result int64
		err := tx.Count(&result).Error
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	if in.ProblemTitle == "" && len(in.TagIds) == 0 && in.Difficulty == 0 {
		// 查询条件为空，缓存总数
		rdb := l.svcCtx.RDB
		key := "cache:problem_count"
		_, err := redis_cache.QueryWithCache(rdb, key, &count, queryCount)
		if err != nil {
			return nil, err
		}
	} else {
		// 直接查询总数
		pCount, err := queryCount()
		if err != nil {
			return nil, err
		}
		count = *pCount
	}

	// 分页查询
	offset := (in.PageNum - 1) * in.PageSize
	var problems []models.Problem
	err := tx.Select("problems.id").Offset(int(offset)).Limit(int(in.PageSize)).Find(&problems).Error
	if err != nil {
		return nil, err
	}

	// 统计结果
	problemIds := make([]int64, len(problems))
	for i, problem := range problems {
		problemIds[i] = problem.ID
	}

	return &problemset.SearchProblemResp{
		ProblemIds: problemIds,
		Count:      count,
	}, nil
}

// IncrByProblemCountCache 题目总数缓存自增
func IncrByProblemCountCache(rdb redis.UniversalClient, value int64) error {
	key := "cache:problem_count"
	return redis_cache.CacheIncrBy(rdb, key, value)
}
