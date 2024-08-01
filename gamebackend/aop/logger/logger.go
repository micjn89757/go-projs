/*
使用zap logger
*/

package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



var (
	Logger *zap.Logger

	loggerOnce sync.Once
)

func InitLogger() {
	loggerOnce.Do(func() {
		if Logger == nil {
			Logger = newLogger("./log/test.log")
		}
	})
}


func newLogger(filePath string) *zap.Logger {
	writeSyncer := getLogWriter(filePath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// 如何写入日志
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder  // 修改时间编码器，格式为RFC3339 YYYY-MM-DDTHH:mm:ss
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志文件中使用大写字母记录日志级别
	return zapcore.NewJSONEncoder(encoderConfig) 
}


// 日志写到哪里
func getLogWriter(filePath string) zapcore.WriteSyncer {
	file ,err := os.OpenFile(filePath, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "log file path not found: err", err)
		os.Exit(1)
	}

	writer :=  io.MultiWriter(os.Stderr, file)
	return zapcore.AddSync(writer)
}