package models

type Difficulty uint32

const (
	Easy Difficulty = iota + 1
	Medium
	Hard
)
