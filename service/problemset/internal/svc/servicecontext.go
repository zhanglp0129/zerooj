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
	RW                 snowflake.WorkerInterface
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
	machineId := c.Database.DataSource[0].MachineId
	snowflakeKey := fmt.Sprintf("/snowflake/problemset/%d", machineId)
	rw, err := redis_snowflake.NewRedisWorkerNoLock(rdb, snowflakeKey, snowflake.NewDefaultConfigWithStartTime(constant.StartTime), machineId)
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
