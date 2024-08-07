package models

type Follow struct {
	FollowerID uint `gorm:"primaryKey;comment:关注者"`
	FollowedID uint `gorm:"primaryKey;comment:被关注者"`
	Follower   User
	Followed   User
}
