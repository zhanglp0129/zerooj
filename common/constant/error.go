package constant

import "fmt"

const (
	RedisPingError           = "ping redis error"
	NotSupportDatabaseType   = "not support the database type"
	NotSupportDatabaseDriver = "not support the database driver"
	NumberOfDataSourceError  = "number of data source error"
	DatabaseMachineIdError   = "database machine id error"
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

// OperationInCoolPeriod 处于冷却期
type OperationInCoolPeriod struct{}

func (OperationInCoolPeriod) Error() string {
	return "当前操作处于冷却期"
}

type NewEqualsOldError struct {
	Thing string
}

func (e NewEqualsOldError) Error() string {
	return fmt.Sprintf("新%s与旧%s相同", e.Thing, e.Thing)
}
