package interfaces

import logging "github.com/sudzekai/golang-logging"

type LoggerFactoryInterface interface {
	NewLogger(category string) LoggerInterface

	SetMinLevel(level logging.LogLevel)
	GetMinLevel() logging.LogLevel

	EnableCliSymbol()
	DisableCliSymbol()
}
