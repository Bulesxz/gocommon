package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

func Sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}
func Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func With(fields ...zapcore.Field) *zap.SugaredLogger {
	return logger.With(fields...).Sugar()
}
