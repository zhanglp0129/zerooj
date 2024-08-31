package svc

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"zerooj/service/mail/internal/config"
)

type ServiceContext struct {
	Config config.Config
	RDB    redis.UniversalClient
	Lim    *redis_rate.Limiter
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb, err := init.InitRedis(c.RedisClient)
	if err != nil {
		panic(err)
	}

	lim := redis_rate.NewLimiter(rdb)

	return &ServiceContext{
		Config: c,
		RDB:    rdb,
		Lim:    lim,
	}
}
