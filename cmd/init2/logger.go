package init2

import "shop-test/pkg/log"

func NewLogger(outputFile string, debug bool, encoding string, name string) log.ILogger {
	return log.NewDevelopmentLogger(outputFile, debug, encoding, name)
}





