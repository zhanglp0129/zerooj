package common

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
	"zerooj/common/constant"
)

// InCoolPeriod 判断是否处于冷却期。处于冷却期返回constant.OperationInCoolPeriod；
// 不处于冷却期，返回nil
func InCoolPeriod(rdb redis.UniversalClient, key string) error {
	_, err := rdb.Get(context.Background(), key).Result()
	if err == nil {
		// 处于冷却期
		return constant.OperationInCoolPeriod
	} else if !errors.Is(err, redis.Nil) {
		// 其他错误
		return err
	}
	return nil
}

// AddCoolPeriod 加入冷却期
func AddCoolPeriod(rdb redis.UniversalClient, key string, period time.Duration) error {
	return rdb.Set(context.Background(), key, "1", period).Err()
}

// AddCoolPeriodToPipe 把 加入冷却期 的操作放到pipe中
func AddCoolPeriodToPipe(pipe redis.Pipeliner, key string, period time.Duration) {
	pipe.Set(context.Background(), key, "1", period)
}
