package log

import (
	"os"
	"path/filepath"

	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(outputLog string, name ...string) ILogger {
	if _, err := os.Stat(outputLog); os.IsNotExist(err) {
		_, err = create(outputLog)
		if err != nil {
			log.Panic(err)
		}
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{outputLog, "stdout"}

	logger, err := cfg.Build()
	if err != nil {
		log.Panic(err)
	}

	slogger := logger.Sugar()

	if len(name) > 0 {
		return slogger.Named(name[0])
	} else {
		return slogger
	}
}

func NewDevelopmentLogger(outputLog string, debug bool, encoding string, name ...string) ILogger {

	cfg := zap.NewDevelopmentConfig()

	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	if debug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	cfg.Encoding = encoding
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	cfg.OutputPaths = []string{"stdout"}

	if debug {
		if _, err := os.Stat(outputLog); os.IsNotExist(err) {
			_, err = create(outputLog)
			if err != nil {
				log.Panic(err)
			}
		}

		cfg.OutputPaths = append(cfg.OutputPaths, outputLog)
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Panic(err)
	}

	slogger := logger.Sugar()

	if len(name) > 0 {
		return slogger.Named(name[0])
	} else {
		return slogger
	}
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
