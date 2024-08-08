/*
定义logger接口
*/

package logger

type Logger interface {
	Debug(msg string, v ...any)
	Info(msg string, v ...any)
	Warn(msg string, v ...any)
	Error(msg string, v ...any)
	Panic(msg string, v ...any)
	Fatal(msg string, v ...any)
}


type defaultLogger struct {
	level Level
}

func Debug() {

}

func Info() {

}

func Warn() {

}

func Error() {

}

func Panic() {

}

func Fatal() {

}