/*
定义logger接口
*/

package logger

import "io"

type LoggerType = byte

const (
	SugarLogger LoggerType = iota 
)

type Logger interface {
	Debug(msg string, v ...any)
	Info(msg string, v ...any)
	Warn(msg string, v ...any)
	Error(msg string, v ...any)
	Panic(msg string, v ...any)
	Fatal(msg string, v ...any)
}


func NewLogger(loggerType LoggerType,  writer io.Writer) (Logger, error) {
	switch loggerType {
	case SugarLogger:
		return NewSugar(writer)
	default:
		panic("unsupported io type")
	}
}


// type defaultLogger struct {
// 	level Level
// }

// func Debug() {

// }

// func Info() {

// }

// func Warn() {

// }

// func Error() {

// }

// func Panic() {

// }

// func Fatal() {

// }