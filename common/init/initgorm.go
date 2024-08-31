package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zerooj/common/constant"
)

type DatabaseConfig struct {
	Driver    string
	DSN       string
	MachineId int64
}

func InitGorm(c DatabaseConfig) (*gorm.DB, error) {
	// 检查配置
	if c.Driver != "mysql" {
		return nil, constant.NotSupportDatabaseDriver
	}
	if c.MachineId != 0 {
		return nil, constant.DatabaseMachineIdError
	}

	db, err := gorm.Open(mysql.Open(c.DSN))
	return db, err
}
