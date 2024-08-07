package models

import (
	"gorm.io/gorm"
	"zerooj/common"
)

// User 用户基本信息，用于登录
type User struct {
	gorm.Model
	Username   string            `gorm:"size:32;not null;unique"`
	Password   string            `gorm:"size:64;not null;comment:加密后的密码"`
	Email      string            `gorm:"size:128;not null;unique"`
	Permission common.Permission `gorm:"not null;comment:用户权限"`
	Profile    UserProfile
	Websites   []UserPersonalWebsite
	Skills     []UserSkill
	Followings []Follow `gorm:"foreignKey:FollowerID"` // 关注
	Fans       []Follow `gorm:"foreignKey:FollowedID"` // 粉丝
}
