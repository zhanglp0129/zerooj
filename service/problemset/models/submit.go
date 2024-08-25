package models

import (
	"time"
	"zerooj/common/models"
)

type Submit struct {
	models.Model
	UserID            int64 `gorm:"index"`
	ProblemID         int64 `gorm:"index;not null"`
	Problem           *Problem
	LanguageID        int64
	Language          *Language
	Result            models.SubmitResult
	CodeObjectName    string        `gorm:"size:100;comment:提交的代码所在的minio对象名称"`
	RunTime           time.Duration `gorm:"comment:执行时间，单位为纳秒"`
	MemoryConsumption models.Memory `gorm:"comment:消耗内存，单位为字节"`
}
