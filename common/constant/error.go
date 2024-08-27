package constant

const (
	NumberOfDataSourceError      ZeroOJError = "数据源数量错误"
	DatabaseMachineIdError       ZeroOJError = "数据库机器码错误"
	FollowYourselfError          ZeroOJError = "不能关注自己"
	NotSupportDatabaseType       ZeroOJError = "不支持的数据库类型"
	NotSupportDatabaseDriver     ZeroOJError = "不支持的数据库驱动"
	UsernamePasswordError        ZeroOJError = "用户名或密码错误"
	EmailPasswordError           ZeroOJError = "邮箱或密码错误"
	MailCheckCodeError           ZeroOJError = "邮箱验证码错误"
	OldMailCheckCodeError        ZeroOJError = "旧邮箱验证码错误"
	NewMailCheckCodeError        ZeroOJError = "新邮箱验证码错误"
	OldPasswordError             ZeroOJError = "旧密码错误"
	OperationInCoolPeriod        ZeroOJError = "当前操作处于冷却期"
	NewPasswordEqualsOldError    ZeroOJError = "新密码与旧密码相同"
	NewEmailEqualsOldError       ZeroOJError = "新邮箱与旧邮箱相同"
	NewUsernameEqualsOldError    ZeroOJError = "新用户名与旧用户名相同"
	FrequentVisitError           ZeroOJError = "您的访问过于频繁，请稍后再试"
	InsufficientPermissionsError ZeroOJError = "权限不足"
	SkillQuantityExceedsLimit    ZeroOJError = "技能数量超出限制"
	WebsiteQuantityExceedsLimit  ZeroOJError = "个人网站数量超出限制"
	NotSupportStorageType        ZeroOJError = "不支持的存储类型"
)

type ZeroOJError string

func (e ZeroOJError) Error() string {
	return string(e)
}
