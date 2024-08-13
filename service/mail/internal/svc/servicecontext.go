package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"zerooj/common/constant"
	"zerooj/service/mail/internal/config"
)

type ServiceContext struct {
	Config config.Config
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.MustNewRedis(c.RedisClient)
	if !rdb.Ping() {
		panic(errors.New(constant.RedisPingError))
	}

	return &ServiceContext{
		Config: c,
		RDB:    rdb,
	}
}
