package models

type Memory int64

const (
	B Memory = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)
