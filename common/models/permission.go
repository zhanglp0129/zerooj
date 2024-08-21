package models

// Permission 用户权限
type Permission uint32

const (
	AdminPermission   Permission = 0
	SupportPermission Permission = 1
	DefaultPermission Permission = 6
	MutePermission    Permission = 7
	BannedPermission  Permission = 8
)

// CanAdmin 是否为管理员
func (p Permission) CanAdmin() bool {
	return p == AdminPermission
}

// CanSupport 是否为客服
func (p Permission) CanSupport() bool {
	return p <= SupportPermission
}

// CanComment 是否能评论
func (p Permission) CanComment() bool {
	return p < MutePermission
}

// CanLogin 是否能登录
func (p Permission) CanLogin() bool {
	return p < BannedPermission
}
