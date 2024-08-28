package models

type SubmitResult uint32

const (
	Waiting SubmitResult = iota + 1
	Judging
	Accepted
	WrongAnswer
	RuntimeError
	CompileLimitExceed
	CompileError
	PresentationError
	TimeLimitExceed
	MemoryLimitExceed
	OutputLimitExceed
	SystemError
)
