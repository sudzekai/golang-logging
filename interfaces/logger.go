package interfaces

import logging "github.com/sudzekai/golang-logging"

type LoggerInterface interface {
	Log(logging.LogEntry)

	LogDebug(format string, args ...any)
	LogInformation(format string, args ...any)
	LogWarning(format string, args ...any)
	LogError(format string, args ...any)
	LogCritical(format string, args ...any)

	IsEnabled(level logging.LogLevel) bool
}
