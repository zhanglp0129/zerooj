package models

type ProblemStatus uint32

const (
	NotStart = iota + 1
	Tried
	Resolved
)
