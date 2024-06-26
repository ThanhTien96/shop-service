package log

type (
	ILogger interface {
		DPanic(args ...interface{})
		DPanicf(template string, args ...interface{})
		Debug(args ...interface{})
		Debugf(template string, args ...interface{})
		Error(args ...interface{})
		Errorf(template string, args ...interface{})
		Fatal(args ...interface{})
		Fatalf(template string, args ...interface{})
		Info(args ...interface{})
		Infof(template string, args ...interface{})
		Panic(args ...interface{})
		Panicf(template string, args ...interface{})
		Sync() error
		Warn(args ...interface{})
		Warnf(template string, args ...interface{})
	}
)