package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"zerooj/common"
	"zerooj/common/constant"
	"zerooj/service/user/internal/config"
	"zerooj/service/user/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := common.InitGorm(c.Database)
	if err != nil {
		panic(err)
	}
	// 建表
	err = db.AutoMigrate(
		&models.City{},
		&models.User{},
		&models.UserProfile{},
		&models.PersonalWebsite{},
		&models.Skill{},
		&models.Follow{},
	)
	if err != nil {
		panic(err)
	}

	rdb := redis.MustNewRedis(c.RedisClient)
	if !rdb.Ping() {
		panic(errors.New(constant.RedisPingError))
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    rdb,
	}
}
