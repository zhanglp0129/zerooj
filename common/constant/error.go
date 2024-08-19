package constant

import "fmt"

const (
	RedisPingError           = "ping redis error"
	NotSupportDatabaseType   = "not support the database type"
	NotSupportDatabaseDriver = "not support the database driver"
	NumberOfDataSourceError  = "number of data source error"
	DatabaseMachineIdError   = "database machine id error"
)

// InputDataError 输入数据错误
type InputDataError struct {
	Thing string
}

func (e InputDataError) Error() string {
	return fmt.Sprintf("%s错误", e.Thing)
}

// OperationInCoolPeriod 处于冷却期
type OperationInCoolPeriod struct{}

func (OperationInCoolPeriod) Error() string {
	return "当前操作处于冷却期"
}

// NewEqualsOldError 新旧相等
type NewEqualsOldError struct {
	Thing string
}

func (e NewEqualsOldError) Error() string {
	return fmt.Sprintf("新%s与旧%s相同", e.Thing, e.Thing)
}
