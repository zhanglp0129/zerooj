package constant

import "time"

var (
	// Location 时区，东八区
	Location, _ = time.LoadLocation("Asia/Shanghai")
	// StartTime 起始时间，2024-8-10 00:00:00
	StartTime = time.Date(2024, time.August, 10, 0, 0, 0, 0, Location)
)

const (
	// MailCheckCodeExpirationTime 邮箱验证码过期时间
	MailCheckCodeExpirationTime = 5 * time.Minute
)
