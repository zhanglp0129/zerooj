package models

import "zerooj/common/models"

type JudgeData struct {
	models.Model
	ProblemID        int64  `gorm:"index;not null"`
	InputObjectName  string `gorm:"size:100;comment:输入数据在minio中的对象名"`
	OutputObjectName string `gorm:"size:100;comment:输出数据在minio中的对象名"`
}
