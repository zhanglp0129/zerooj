package taglogic

import (
	"context"
	"gorm.io/gorm"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加标签
func (l *AddTagLogic) AddTag(in *problemset.AddTagReq) (*problemset.AddTagResp, error) {
	var err error
	t := models.Tag{
		Name: in.Name,
	}
	t.ID, err = l.svcCtx.RW.GenerateId()
	if err != nil {
		return nil, err
	}

	db := l.svcCtx.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		// 查询数据
		err = db.Create(&t).Error
		if err != nil {
			return err
		}

		// 删除缓存
		return DeleteAllTagsCache(l.svcCtx.RDB)
	})

	return &problemset.AddTagResp{
		Id: t.ID,
	}, nil
}
