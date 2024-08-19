package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// FinishMailCheckCode 使用完了邮箱验证码，应该删除，失败了也不要紧
func FinishMailCheckCode(rdb redis.UniversalClient, key string) {
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
