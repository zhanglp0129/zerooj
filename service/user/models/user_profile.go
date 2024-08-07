package models

import (
	"gorm.io/gorm"
	"time"
)

// UserProfile 用户附加信息
type UserProfile struct {
	gorm.Model
	UserID               uint      `gorm:"index;not null;unique"`
	Nickname             string    `gorm:"size:30;comment:昵称"`
	AvatarURL            string    `gorm:"size:128;comment:头像URL"`
	Gender               uint8     `gorm:"comment:性别，0不愿意透露，1男，2女"`
	Birthday             time.Time `gorm:"comment:出生日期"`
	PersonalIntroduction string    `gorm:"size:200;comment:个人介绍"`
	CityID               uint
	City                 City
}

// UserPersonalWebsite 用户个人网站
type UserPersonalWebsite struct {
	gorm.Model
	UserID uint   `gorm:"index"`
	URL    string `gorm:"size:128;comment:个人网站URL"`
}

// UserSkill 用户技能
type UserSkill struct {
	gorm.Model
	UserID uint `gorm:"index"`
	Skill  string
}
