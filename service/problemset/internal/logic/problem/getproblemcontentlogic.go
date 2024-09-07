package problemlogic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_cache"
	"zerooj/common/models/storagetype"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemContentLogic {
	return &GetProblemContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取问题信息，可缓存
func (l *GetProblemContentLogic) GetProblemContent(in *problemset.GetProblemContentReq, stream problemset.Problem_GetProblemContentServer) error {
	db, rdb := l.svcCtx.DB, l.svcCtx.RDB
	key := fmt.Sprintf("cache:problem_content:%d", in.GetProblemId())

	var model problemset.ProblemInfo
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*problemset.ProblemInfo, error) {
		// 查询数据库
		var p models.Problem
		p.ID = in.ProblemId
		err := db.Take(&p).Error
		if err != nil {
			return nil, err
		}

		// 判断存储类型，获取原始数据
		result := problemset.ProblemInfo{
			ProblemId:  p.ID,
			Title:      p.Title,
			Difficulty: uint32(p.Difficulty),
		}
		// 题目描述
		byts, err := storagetype.GetContent(p.DescriptionStorageType, []byte(p.Description))
		if err != nil {
			return nil, err
		}
		result.Description = string(byts)
		// 输入描述
		byts, err = storagetype.GetContent(p.InputDescriptionStorageType, []byte(p.InputDescription))
		if err != nil {
			return nil, err
		}
		result.InputDescription = string(byts)
		// 输出描述
		byts, err = storagetype.GetContent(p.OutputDescriptionStorageType, []byte(p.OutputDescription))
		if err != nil {
			return nil, err
		}
		result.OutputDescription = string(byts)
		// 规模描述
		byts, err = storagetype.GetContent(p.ScaleDescriptionStorageType, []byte(p.ScaleDescription))
		if err != nil {
			return nil, err
		}
		result.ScaleDescription = string(byts)

		return &result, nil
	})
	if err != nil {
		return err
	}

	// 发送响应
	return stream.Send(&problemset.GetProblemContentResp{
		Problem: &model,
	})
}

// DeleteProblemContentCache 删除题目信息缓存
func DeleteProblemContentCache(rdb redis.UniversalClient, problemId int64) error {
	key := fmt.Sprintf("cache:problem_content:%d", problemId)
	return redis_cache.DeleteCache(rdb, key)
}
