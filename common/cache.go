package common

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math/rand"
	"reflect"
)

// GetCacheExpireSeconds 获取缓存过期秒数，55-65秒
func GetCacheExpireSeconds() int {
	return 55 + rand.Int()%11
}

type QueryFunc[T any] func() (T, error)

// WithCache 带着缓存查询数据。
// model，传出参数，数据模型；
// query，查询函数
func WithCache[T any](rdb *redis.Redis, key string, model *T, query QueryFunc[*T]) error {
	rVal := reflect.ValueOf(model).Elem()
	rTyp := rVal.Type()

	exists, err := rdb.Exists(key)
	if err != nil {
		return err
	}
	if exists {
		// 命中缓存
		all, err := rdb.Hgetall(key)
		if err != nil {
			return err
		}
		// 赋值
		for i := 0; i < rTyp.NumField(); i++ {
			if !rTyp.Field(i).IsExported() {
				continue
			}
			fieldName := rTyp.Field(i).Name
			stringValue := all[fieldName]

			rVal.Field(i).SetString(stringValue)
		}
		// 重设过期时间
		if err = rdb.Expire(key, GetCacheExpireSeconds()); err != nil {
			return err
		}
	}

	// 未命中缓存
	data, err := query()
	if err != nil {
		return err
	}
	if data == nil {
		// 没有数据
		return nil
	}
	*model = *data
	mp := make(map[string]string)
	// 写入缓存
	for i := 0; i < rTyp.NumField(); i++ {
		if !rTyp.Field(i).IsExported() {
			continue
		}
		fieldName := rTyp.Field(i).Name
		fieldValue := reflect.ValueOf(data).Elem().Field(i)
		mp[fieldName] = fieldValue.String()
	}
	if err = rdb.Hmset(key, mp); err != nil {
		return err
	}
	// 设置过期时间
	if err = rdb.Expire(key, GetCacheExpireSeconds()); err != nil {
		return err
	}

	return nil
}

func DeleteCache(rdb *redis.Redis, key string) error {
	_, err := rdb.Del(key)
	return err
}
