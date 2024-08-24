package models

import (
	"zerooj/common/models"
)

type Example struct {
	models.Model
	ProblemID   int64  `gorm:"index;not null"`
	Input       string `gorm:"size:10000;comment:输入数据"`
	Output      string `gorm:"size:10000;comment:输出数据"`
	Explanation string `gorm:"size:1000;comment:解释"`
}
