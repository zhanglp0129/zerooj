package models

import (
	"github.com/zhanglp0129/snowflake"
	"gorm.io/gorm"
	"time"
	"zerooj/common/constant"
)

type Model struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	worker *snowflake.Worker
)

func init() {
	// 创建雪花算法节点，机器码为0
	sc := snowflake.DefaultConfig
	sc.SetStartTime(constant.StartTime)
	var err error
	worker, err = snowflake.NewWorker(snowflake.DefaultConfig, 0)
	if err != nil {
		panic(err)
	}
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID, err = worker.GenerateId()
	return
}
