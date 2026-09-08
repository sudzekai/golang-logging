package logging

type LogLevel int

const (
	Debug LogLevel = iota
	Information
	Warning
	Error
	Critical
)
