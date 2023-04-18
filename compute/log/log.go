package log

import (
	"os"
	"strings"
	"time"

	cfg "compute/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

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

const (
	debugLevel = "debug"
	infoLevel  = "info"
	warnLevel  = "warn"
	errorLevel = "error"
)

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case debugLevel:
		return zapcore.DebugLevel
	case infoLevel:
		return zapcore.InfoLevel
	case errorLevel:
		return zapcore.ErrorLevel
	case warnLevel:
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
	fe := fileEncoder()
	std := stdEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(std, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(fe, getLogWriter(), level),
	)
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
