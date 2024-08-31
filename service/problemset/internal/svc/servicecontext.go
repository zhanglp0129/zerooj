package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/snowflake"
	"gorm.io/gorm"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/problemset/internal/config"
	"zerooj/service/problemset/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    redis.UniversalClient
	RW     snowflake.WorkerInterface
	MC     *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := common.InitGorm(c.Database)
	if err != nil {
		panic(err)
	}
	// 建表
	err = db.AutoMigrate(
		&models.Problem{},
		&models.Example{},
		&models.Hint{},
		&models.JudgeData{},
		&models.Tag{},
		&models.Language{},
		&models.ProblemLanguage{},
		&models.Submit{},
	)
	if err != nil {
		panic(err)
	}

	// 初始化redis
	rdb, err := common.InitRedis(c.RedisClient)
	if err != nil {
		panic(err)
	}

	// 创建雪花算法节点
	machineId := c.Database.MachineId
	rw, err := snowflake.NewWorker(snowflake.NewDefaultConfigWithStartTime(constant.StartTime), machineId)
	if err != nil {
		panic(err)
	}

	// 创建minio客户端
	mc, err := common.InitMinio(c.Minio)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    rdb,
		RW:     rw,
		MC:     mc,
	}
}
