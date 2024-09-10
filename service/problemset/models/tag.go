package models

import "zerooj/common/models"

type Tag struct {
	models.Model
	Name string `gorm:"size:20;unique"`
}
