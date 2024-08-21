package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// FinishMailCheck 完成邮箱验证，应该删除邮箱验证码。失败了也不要紧，可以并发执行
func FinishMailCheck(rdb redis.UniversalClient, key string) {
	// 失败后重试
	for i := 0; i < 3; i++ {
		err := rdb.Del(context.Background(), key).Err()
		if err == nil {
			break
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}
