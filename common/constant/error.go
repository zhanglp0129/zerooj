package constant

const (
	RedisPingError = "ping redis error"
)

type MailCheckCodeError struct{}

func (e MailCheckCodeError) Error() string {
	return "邮件验证码错误"
}
