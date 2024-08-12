package constant

const (
	RedisPingError = "ping redis error"
)

// MailCheckCodeError 邮件验证码错误
type MailCheckCodeError struct{}

func (MailCheckCodeError) Error() string {
	return "邮件验证码错误"
}

// UsernameEmailPasswordError 用户名、邮箱或密码错误
type UsernameEmailPasswordError struct{}

func (UsernameEmailPasswordError) Error() string {
	return "用户名、邮箱或密码错误"
}
