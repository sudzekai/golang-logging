package logging

import (
	"time"
)

type LogEntry struct {
	TimeStamp time.Time
	Level     LogLevel
	Category  string
	Message   string
}
