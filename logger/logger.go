package logger

import "fmt"

type (
	Logger interface {
		Debug(fmtStr string, args ...interface{})
		Info(fmtStr string, args ...interface{})
		Warn(fmtStr string, args ...interface{})
		Error(fmtStr string, args ...interface{})
	}

	defLogger struct{}
)

func (d defLogger) Debug(fmtStr string, args ...interface{}) {
	fmt.Println("Debug", fmt.Sprintf(fmtStr, args...))
}

func (d defLogger) Info(fmtStr string, args ...interface{}) {
	fmt.Println("Info", fmt.Sprintf(fmtStr, args...))
}

func (d defLogger) Warn(fmtStr string, args ...interface{}) {
	fmt.Println("Warn", fmt.Sprintf(fmtStr, args...))
}

func (d defLogger) Error(fmtStr string, args ...interface{}) {
	fmt.Println("Error", fmt.Sprintf(fmtStr, args...))
}

var logger Logger

func init() {
	logger = &defLogger{}
}

func GetLogger() Logger {
	return logger
}

func SetLogger(l Logger) {
	logger = l
}
