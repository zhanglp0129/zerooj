package svc

import (
	"github.com/redis/go-redis/v9"
	"zerooj/common"
	"zerooj/service/authenticate/internal/config"
)

type ServiceContext struct {
	Config config.Config
	RDB    redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化redis
	rdb, err := common.InitRedis(c.RedisClient)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		RDB:    rdb,
	}
}
