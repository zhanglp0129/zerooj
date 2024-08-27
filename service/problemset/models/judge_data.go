package models

import (
	"zerooj/common/models"
	"zerooj/common/models/storagetype"
)

type JudgeData struct {
	models.Model
	ProblemID         int64                   `gorm:"index;not null"`
	InputStorageType  storagetype.StorageType `gorm:"comment:输入存储类型"`
	Input             string                  `gorm:"size:255;comment:输入数据"`
	OutputStorageType storagetype.StorageType `gorm:"comment:输出存储类型"`
	Output            string                  `gorm:"size:255;comment:输出数据"`
}
