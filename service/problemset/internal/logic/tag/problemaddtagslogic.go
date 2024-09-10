package taglogic

import (
	"context"
	"gorm.io/gorm"
	"zerooj/common/constant"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemAddTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProblemAddTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemAddTagsLogic {
	return &ProblemAddTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给题目添加标签，最多5个
func (l *ProblemAddTagsLogic) ProblemAddTags(in *problemset.ProblemAddTagsReq) (*problemset.ProblemAddTagsResp, error) {
	// 判断标签数量
	db := l.svcCtx.DB
	var p models.Problem
	p.ID = in.ProblemId
	count := db.Model(&p).Association("Tags").Count()
	if int(count)+len(in.TagIds) > 5 {
		return nil, constant.ProblemTagQuantityExceedsLimit
	}

	// 添加标签
	p = models.Problem{}
	p.ID = in.ProblemId
	tags := make([]models.Tag, len(in.TagIds))
	for i := range tags {
		tags[i].ID = in.TagIds[i]
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Model(&p).Association("Tags").Append(&tags)
		if err != nil {
			return err
		}

		// 删除缓存
		return DeleteProblemTagsCache(l.svcCtx.RDB, in.ProblemId)
	})
	if err != nil {
		return nil, err
	}

	return &problemset.ProblemAddTagsResp{}, nil
}
