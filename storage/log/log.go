package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	cfg "storage/config"
	"strings"
	"time"
)

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Cfg.FileName,
		MaxSize:    cfg.Cfg.MaxSize,
		MaxBackups: cfg.Cfg.MaxBackUps,
		MaxAge:     cfg.Cfg.MaxAge,
		Compress:   cfg.Cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func fileEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	// 时间格式
	ec.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	ec.EncodeLevel = zapcore.LowercaseLevelEncoder
	if cfg.Cfg.Encoding == "json" {
		return zapcore.NewJSONEncoder(ec)
	} else {
		return zapcore.NewConsoleEncoder(ec)
	}
}

func stdEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	ec.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	if cfg.Cfg.Encoding == "json" {
		return zapcore.NewJSONEncoder(ec)
	} else {
		return zapcore.NewConsoleEncoder(ec)
	}
}

var log *zap.Logger

const (
	DebugLevel = "debug"
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
	level := getLevel(cfg.Cfg.Level)
	zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DebugLevel
	})(level)
	var core zapcore.Core
	if cfg.Cfg.WriteConsole {
		std := stdEncoder()
		fe := fileEncoder()
		core = zapcore.NewTee(
			zapcore.NewCore(std, zapcore.AddSync(os.Stdout), level),
			zapcore.NewCore(fe, getLogWriter(), level),
		)
	} else {
		fe := fileEncoder()
		core = zapcore.NewCore(fe, getLogWriter(), level)
	}
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}
func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zapcore.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zapcore.Field) {
	log.Error(message, fields...)
}

func Panic(message string, fields ...zapcore.Field) {
	log.Panic(message, fields...)
}
