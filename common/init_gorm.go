package common

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

type GormConfig struct {
	Mysql struct {
		DataSource string
	}
	Log struct {
		Level string
		File  string
	}
}

func InitGorm(c GormConfig) (*gorm.DB, error) {
	// 实现日志配置
	var logLevel logger.LogLevel
	switch c.Log.Level {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		return nil, errors.New("不支持的日志级别")
	}
	var logOut io.Writer
	var colorful bool
	if c.Log.File == "" {
		logOut = os.Stdout
		colorful = true
	} else {
		var err error
		logOut, err = os.OpenFile(c.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		colorful = false
	}

	gormLog := logger.New(log.New(logOut, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logLevel,
		IgnoreRecordNotFoundError: false,
		Colorful:                  colorful,
	})

	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: gormLog,
	})
	return db, err
}
