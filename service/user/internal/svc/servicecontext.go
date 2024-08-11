package svc

import (
	"gorm.io/gorm"
	"zerooj/common"
	"zerooj/service/user/internal/config"
	"zerooj/service/user/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := common.InitGorm(c.Gorm)
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

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
