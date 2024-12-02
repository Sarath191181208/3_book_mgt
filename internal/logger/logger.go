package logger

import "fmt"

type Logger interface {
	Debug(s ...any)
	Info(s ...any)
	Error(err error, s ...any)
}

type SysOutLogger struct{}

func (l SysOutLogger) Debug(s ...any) {
	fmt.Println(s...)
}

func (l SysOutLogger) Info(s ...any) {
	fmt.Println(s...)
}

func (l SysOutLogger) Error(err error, s ...any) {
	fmt.Println(s...)
	fmt.Println(err)
}

func NewSysOutLogger() Logger {
	return SysOutLogger{}
}
