package models

import (
	"zerooj/common/models"
)

type City struct {
	models.Model
	Name string `gorm:"size:10;not null;unique;comment:城市名称"`
}
