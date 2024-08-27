package models

import (
	"time"
	"zerooj/common/models"
)

type Language struct {
	models.Model
	Name    string `gorm:"size:20;comment:语言名称"`
	Version string `gorm:"size:20;comment:语言版本"`
}

type ProblemLanguage struct {
	ProblemID   int64         `gorm:"primaryKey"`
	LanguageID  int64         `gorm:"primaryKey"`
	TimeLimit   time.Duration `gorm:"comment:时间限制，单位为纳秒"`
	MemoryLimit models.Memory `gorm:"comment:内存限制，单位为字节"`
}
