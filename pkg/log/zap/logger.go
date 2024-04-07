package zap

import (
	"shop-test/pkg/log"

	"go.uber.org/zap"
)

func New(name ...string) log.ILogger {
	return NewLogger()
}

func NewLogger(name ...string) log.ILogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	if len(name) > 0 {
		return slogger.Named(name[0])
	} else {
		return slogger
	}
}

func NewDevelopment(name ...string) log.ILogger {
	logger, _ := zap.NewDevelopment()
	slogger := logger.Sugar()

	if len(name) > 0 {
		return slogger.Named(name[0])
	} else {
		return slogger
	}
}
