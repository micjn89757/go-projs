/*
日志级别
*/

package logger

type Level uint 

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
)

func SetLevel(lv Level) Level {
	if lv < LevelDebug || lv > LevelFatal {
		panic("invalid level")
	}

	return lv
}