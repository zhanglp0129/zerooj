package models

import "zerooj/common/models"

type Hint struct {
	models.Model
	ProblemID int64  `gorm:"index;not null"`
	Content   string `gorm:"size:1000"`
}
