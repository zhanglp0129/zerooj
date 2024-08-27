package models

import (
	"zerooj/common/models"
)

type Problem struct {
	models.Model
	Title                 string            `gorm:"size:20;comment:题目标题"`
	Difficulty            models.Difficulty `gorm:"comment:难度，1简单，2中等，3困难"`
	DescriptionObjectName string            `gorm:"size:100;comment:题目描述，在minio的对象名"`
	InputDescription      string            `gorm:"size:1000;comment:输入描述"`
	OutputDescription     string            `gorm:"size:1000;comment:输出描述"`
	ScaleDescription      string            `gorm:"size:1000;comment:数据规模说明"`
	Examples              []Example
	Hints                 []Hint
	JudgeData             []JudgeData
	Tags                  []Tag      `gorm:"many2many:problem_tags"`
	Languages             []Language `gorm:"many2many:problem_languages"`
	ProblemLanguages      []ProblemLanguage
	Submits               []Submit
}
