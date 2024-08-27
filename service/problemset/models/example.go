package models

import (
	"zerooj/common/models"
	"zerooj/common/models/storagetype"
)

type Example struct {
	models.Model
	ProblemID              int64                   `gorm:"index;not null"`
	InputStorageType       storagetype.StorageType `gorm:"comment:输入存储类型"`
	Input                  string                  `gorm:"size:255;comment:输入数据"`
	OutputStorageType      storagetype.StorageType `gorm:"comment:输出存储类型"`
	Output                 string                  `gorm:"size:255;comment:输出数据"`
	ExplanationStorageType storagetype.StorageType `gorm:"comment:解释存储类型"`
	Explanation            string                  `gorm:"size:255;comment:解释"`
}
