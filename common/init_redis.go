package common

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisConf struct {
	Addrs      []string
	ClientName string `json:",optional"`
	DB         int    `json:",default=0"`
	Username   string `json:",optional"`
	Password   string `json:",optional"`
}

func InitRedis(c RedisConf) (redis.UniversalClient, error) {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      c.Addrs,
		ClientName: c.ClientName,
		DB:         c.DB,
		Username:   c.Username,
		Password:   c.Password,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
