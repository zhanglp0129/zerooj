package models

import (
	"zerooj/common/models"
	"zerooj/common/models/storagetype"
)

type Problem struct {
	models.Model
	Title                        string                  `gorm:"size:20;comment:题目标题"`
	Difficulty                   models.Difficulty       `gorm:"comment:难度，1简单，2中等，3困难"`
	DescriptionStorageType       storagetype.StorageType `gorm:"comment:题目描述存储类型"`
	Description                  string                  `gorm:"size:255;comment:题目描述"`
	InputDescriptionStorageType  storagetype.StorageType `gorm:"comment:输入描述存储类型"`
	InputDescription             string                  `gorm:"size:255;comment:输入描述"`
	OutputDescriptionStorageType storagetype.StorageType `gorm:"comment:输出描述存储类型"`
	OutputDescription            string                  `gorm:"size:255;comment:输出描述"`
	ScaleDescriptionStorageType  storagetype.StorageType `gorm:"comment:数据规模说明存储类型"`
	ScaleDescription             string                  `gorm:"size:255;comment:数据规模说明"`
	Examples                     []Example
	Hints                        []Hint
	JudgeData                    []JudgeData
	Tags                         []Tag      `gorm:"many2many:problem_tags"`
	Languages                    []Language `gorm:"many2many:problem_languages"`
	ProblemLanguages             []ProblemLanguage
	Submits                      []Submit
}
