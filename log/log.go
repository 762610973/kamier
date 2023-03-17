package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	// 时间格式
	ec.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	return zapcore.NewConsoleEncoder(ec)
}

var Zlog *zap.Logger

const (
	DebugLevel = "debug"
	// InfoLevel default level
	InfoLevel  = "info"
	ErrorLevel = "error"
	WarnLevel  = "warn"
)

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case WarnLevel:
		return zap.WarnLevel
	default:
		return zapcore.PanicLevel
	}
}

func InitLogger() {
	level := getLevel("info")
	zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DebugLevel
	})(level)
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, getLogWriter(), level),
	)
	Zlog = zap.New(core, zap.AddCaller())
}
