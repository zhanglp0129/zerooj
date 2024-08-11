package models

import (
	"time"
	"zerooj/common/models"
)

// UserProfile 用户附加信息
type UserProfile struct {
	models.Model
	UserID               uint64    `gorm:"index;not null;unique"`
	Nickname             string    `gorm:"size:30;comment:昵称"`
	AvatarURL            string    `gorm:"size:128;comment:头像URL"`
	Gender               Gender    `gorm:"comment:性别，0不愿意透露，1男，2女"`
	Birthday             time.Time `gorm:"comment:出生日期"`
	PersonalIntroduction string    `gorm:"size:200;comment:个人介绍"`
	CityID               uint64
	City                 City
}

type Gender uint32

const (
	NotKnowGender Gender = iota
	Male
	Female
)

// PersonalWebsite 用户个人网站
type PersonalWebsite struct {
	models.Model
	UserID uint64 `gorm:"index;not null"`
	Name   string `gorm:"size:20;commit:网站名称"`
	URL    string `gorm:"size:128;not null;comment:个人网站URL"`
}

// Skill 技能
type Skill struct {
	models.Model
	Name string `gorm:"size:20;unique;not null"`
}