package log

import (
	"io"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger
var logger *zap.Logger

func getLogLevel(level string) zapcore.Level {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	}
	panic("unknown level: " + level)
}

// Init the log module
func Init(logLevel string, ws ...io.Writer) {
	level := getLogLevel(logLevel)
	var encoderCfg = zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05.000 MST"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level))
	for _, w := range ws {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), level))
	}
	tree := zapcore.NewTee(cores...)

	if logger != nil {
		_ = logger.Sync()
	}
	if sugar != nil {
		_ = sugar.Sync()
	}
	logger = zap.New(tree)
	sugar = logger.Sugar()
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
		_ = sugar.Sync()
	}
}

func DebugZ(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func InfoZ(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func WarnZ(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func ErrorZ(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}

func Info(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}

func Error(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}
