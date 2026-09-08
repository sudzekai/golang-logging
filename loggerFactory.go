package logging

import "io"

type LoggerFactory struct {
	writer io.Writer
	level  LogLevel
	cli    bool
}

func NewLoggerFactory(w io.Writer) LoggerFactory {
	return LoggerFactory{
		writer: w,
		level:  Information,
	}
}

func (f *LoggerFactory) NewLogger(category string) Logger {
	return Logger{
		writer:   f.writer,
		category: category,
		level:    f.level,
		cli:      false,
	}
}

func (f *LoggerFactory) SetMinLevel(level LogLevel) {
	f.level = level
}

func (f *LoggerFactory) GetMinLevel() LogLevel {
	return f.level
}

func (f *LoggerFactory) EnableCliSymbol() {
	f.cli = true
}

func (f *LoggerFactory) DisableCliSymbol() {
	f.cli = false
}
