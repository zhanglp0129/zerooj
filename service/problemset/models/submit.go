package models

import (
	"time"
	"zerooj/common/models"
	"zerooj/common/models/storagetype"
)

type Submit struct {
	models.Model
	UserID            int64 `gorm:"index"`
	ProblemID         int64 `gorm:"index;not null"`
	Problem           *Problem
	LanguageID        int64
	Language          *Language
	Result            models.SubmitResult
	CodeStorageType   storagetype.StorageType `gorm:"comment:提交的代码存储类型"`
	Code              string                  `gorm:"size:255;comment:提交的代码"`
	RunTime           time.Duration           `gorm:"comment:执行时间，单位为纳秒"`
	MemoryConsumption models.Memory           `gorm:"comment:消耗内存，单位为字节"`
}
