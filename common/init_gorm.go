package common

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zerooj/common/constant"
)

type DatabaseConfig struct {
	Driver     string
	DataSource string
	MachineId  int64
}

func InitGorm(c DatabaseConfig) (*gorm.DB, error) {
	if c.Driver != "mysql" {
		return nil, errors.New(constant.NotSupportDatabaseDriver)
	}

	db, err := gorm.Open(mysql.Open(c.DataSource))
	return db, err
}
