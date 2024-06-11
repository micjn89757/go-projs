package middleware

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"vote-gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 此中间件用来接收并记录gin框架默认的日志
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		logger := initLogger(start)
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)

	}
}


func initLogger(start time.Time) *zap.Logger{
	// 定制logger
	// zapcore.Core需要三个配置——Encoder、WriteSyncer、LogLevel
	// encoder（编码器）如何写入日志
	// WriterSyncer 指定日志写到那里去
	// LogLevel 哪种级别的日志将被写入
	writeSyncer := getLogWriter(start)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	logger := zap.New(core)
	return logger
}

// getEncoder 
func getEncoder() zapcore.Encoder {
	// 我们使用默认提供的NewJSONEncoder以及预先设置的NewProductionEncoderConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 覆盖默认的时间编码
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写字母记录日志级别	
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 
func getLogWriter(start time.Time) zapcore.WriteSyncer {
	// 将日志写入文件
	filePath := filepath.Join(utils.ProjPath, fmt.Sprintf("log/%d-%d-%d.log", start.Year(), start.Month(), start.Day()))
	// 打开文件
	scr, err := os.OpenFile(filePath, os.O_RDWR | os.O_CREATE, 0755)

	if err != nil {
		fmt.Printf("err: %v", err)
		return zapcore.AddSync(os.Stdout)
	}

	// 利用io.MultiWriter将文件和终端作为两个输出目标
	ws := io.MultiWriter(scr, os.Stdout)

	// 使用zapcore.AddSync 将句柄传进去
	return zapcore.AddSync(ws)
}