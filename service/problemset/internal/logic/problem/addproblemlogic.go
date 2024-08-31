package problemlogic

import (
	"context"
	commonmodels "zerooj/common/models"
	"zerooj/common/models/storagetype"
	"zerooj/service/problemset/models"

	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemLogic {
	return &AddProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加问题
func (l *AddProblemLogic) AddProblem(stream problemset.Problem_AddProblemServer) error {
	in, err := stream.Recv()
	if err != nil {
		return err
	}

	// 构建数据模型
	p := &models.Problem{
		Title:      in.Title,
		Difficulty: commonmodels.Difficulty(in.Difficulty),
	}
	p.ID, err = l.svcCtx.RW.GenerateId()
	if err != nil {
		return err
	}

	// 将数据存储至minio
	mc, bucketName := l.svcCtx.MC, l.svcCtx.Config.Minio.BucketName
	// 题目描述
	value, st, err := storagetype.Storage([]byte(in.GetDescription()), mc, bucketName)
	if err != nil {
		return err
	}
	p.Description, p.DescriptionStorageType = string(value), st
	// 输入描述
	value, st, err = storagetype.Storage([]byte(in.GetInputDescription()), mc, bucketName)
	if err != nil {
		return err
	}
	p.InputDescription, p.InputDescriptionStorageType = string(value), st
	// 输出描述
	value, st, err = storagetype.Storage([]byte(in.GetOutputDescription()), mc, bucketName)
	if err != nil {
		return err
	}
	p.OutputDescription, p.OutputDescriptionStorageType = string(value), st
	// 规模描述
	value, st, err = storagetype.Storage([]byte(in.GetScaleDescription()), mc, bucketName)
	if err != nil {
		return err
	}
	p.ScaleDescription, p.ScaleDescriptionStorageType = string(value), st

	// 插入数据库
	db := l.svcCtx.DB
	err = db.Create(&p).Error
	if err != nil {
		return err
	}

	// 响应
	return stream.SendAndClose(&problemset.AddProblemResp{
		Id: p.ID,
	})
}
