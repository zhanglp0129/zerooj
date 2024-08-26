package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zhanglp0129/redis_snowflake"
	"github.com/zhanglp0129/snowflake"
	"gorm.io/gorm"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/user/internal/config"
	"zerooj/service/user/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    redis.UniversalClient
	RW     snowflake.WorkerInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := common.InitGorm(c.Database)
	if err != nil {
		panic(err)
	}
	// 建表
	err = db.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.PersonalWebsite{},
		&models.Skill{},
		&models.Follow{},
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
	machineId := c.Database.DataSource[0].MachineId
	snowflakeKey := fmt.Sprintf("/snowflake/user/%d", machineId)
	rw, err := redis_snowflake.NewRedisWorkerNoLock(rdb, snowflakeKey, snowflake.NewDefaultConfigWithStartTime(constant.StartTime), machineId)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    rdb,
		RW:     rw,
	}
}
