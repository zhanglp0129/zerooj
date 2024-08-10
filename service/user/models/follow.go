package models

import "time"

type Follow struct {
	FollowerID uint64 `gorm:"primaryKey;comment:关注者"`
	FollowedID uint64 `gorm:"primaryKey;comment:被关注者"`
	Follower   User
	Followed   User
	FollowAt   time.Time
}
