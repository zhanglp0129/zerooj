package common

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zerooj/common/constant"
)

type DatabaseConfig struct {
	Type       string
	DataSource []struct {
		Driver    string
		DSN       string
		MachineId int64
	}
}

func InitGorm(c DatabaseConfig) (*gorm.DB, error) {
	// 检查配置
	if c.Type != "node" {
		return nil, errors.New(constant.NotSupportDatabaseType)
	}
	if len(c.DataSource) != 1 {
		return nil, errors.New(constant.NumberOfDataSourceError)
	}
	dataSource := c.DataSource[0]
	if dataSource.Driver != "mysql" {
		return nil, errors.New(constant.NotSupportDatabaseDriver)
	}
	if dataSource.MachineId != 0 {
		return nil, errors.New(constant.DatabaseMachineIdError)
	}

	db, err := gorm.Open(mysql.Open(dataSource.DSN))
	return db, err
}
