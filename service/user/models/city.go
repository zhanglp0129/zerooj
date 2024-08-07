package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name string `gorm:"size:10;not null;unique;comment:城市名称"`
}
