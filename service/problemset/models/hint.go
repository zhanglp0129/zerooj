package models

import (
	"zerooj/common/models"
	"zerooj/common/models/storagetype"
)

type Hint struct {
	models.Model
	ProblemID          int64                   `gorm:"index;not null"`
	ContentStorageType storagetype.StorageType `gorm:"comment:提示内容的存储类型"`
	Content            string                  `gorm:"size:255;comment:提示内容"`
}
