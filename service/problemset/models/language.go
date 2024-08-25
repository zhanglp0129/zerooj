package models

import "zerooj/common/models"

type Language struct {
	models.Model
	Name    string `gorm:"size:20;comment:语言名称"`
	Version string `gorm:"size:20;comment:语言版本"`
}
