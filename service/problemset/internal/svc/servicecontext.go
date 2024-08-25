package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhanglp0129/redis_snowflake"
	"github.com/zhanglp0129/snowflake"
	"gorm.io/gorm"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/problemset/internal/config"
	"zerooj/service/problemset/models"
	"zerooj/service/user/pb/user"
)

type ServiceContext struct {
	Config             config.Config
	DB                 *gorm.DB
	RDB                redis.UniversalClient
	RW                 *redis_snowflake.RedisWorker
	UserBaseInfoClient user.BaseInfoClient
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
	var machineId int64 = 0
	snowflakeKey := fmt.Sprintf("/snowflake/user/%d", machineId)
	snowflakeLockKey := fmt.Sprintf("/snowflake/user/lock/%d", machineId)
	rw, err := redis_snowflake.NewRedisWorker(rdb, snowflakeKey, snowflakeLockKey, snowflake.NewDefaultConfigWithStartTime(constant.StartTime), machineId)
	if err != nil {
		panic(err)
	}

	// 创建用户基本信息服务客户端
	conn := zrpc.MustNewClient(c.Client.User)
	userBaseInfoClient := user.NewBaseInfoClient(conn.Conn())

	return &ServiceContext{
		Config:             c,
		DB:                 db,
		RDB:                rdb,
		RW:                 rw,
		UserBaseInfoClient: userBaseInfoClient,
	}
}
