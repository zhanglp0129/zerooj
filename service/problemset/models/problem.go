package models

import (
	"time"
	"zerooj/common/models"
)

type Problem struct {
	models.Model
	Title             string            `gorm:"size:20;comment:题目标题"`
	Difficulty        models.Difficulty `gorm:"comment:难度，0简单，1中等，2困难"`
	Description       string            `gorm:"size:10000;comment:题目描述"`
	InputDescription  string            `gorm:"size:1000;comment:输入描述"`
	OutputDescription string            `gorm:"size:1000;comment:输出描述"`
	ScaleDescription  string            `gorm:"size:1000;comment:数据规模说明"`
	TimeLimit         time.Duration     `gorm:"comment:时间限制，单位为纳秒"`
	MemoryLimit       int64             `gorm:"comment:内存限制，单位为字节"`
	Examples          []Example
	Hints             []Hint
	JudgeData         []JudgeData
	Tags              []Tag `gorm:"many2many:problem_tags"`
}