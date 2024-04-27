package logger

import (
	"context"
)

type ILogger interface {
	Panic(msg string, keysAndValues ...any)
	Panicf(format string, args ...any)
	Panicc(ctx context.Context, msg string, keysAndValues ...any)
	Paniccf(ctx context.Context, format string, args ...any)

	Error(msg string, keysAndValues ...any)
	Errorf(format string, args ...any)
	Errorc(ctx context.Context, msg string, keysAndValues ...any)
	Errorcf(ctx context.Context, format string, args ...any)

	Info(msg string, keysAndValues ...any)
	Infof(format string, args ...any)
	Infoc(ctx context.Context, msg string, keysAndValues ...any)
	Infocf(ctx context.Context, format string, args ...any)

	Debug(msg string, keysAndValues ...any)
	Debugf(format string, args ...any)
	Debugc(ctx context.Context, msg string, keysAndValues ...any)
	Debugcf(ctx context.Context, format string, args ...any)
}

var log ILogger

func SetLogger(l ILogger) {
	log = l
}

func Panic(msg string, keysAndValues ...any) {
	log.Panic(msg, keysAndValues...)
}

func Panicf(template string, args ...any) {
	log.Panicf(template, args...)
}

func Error(msg string, keysAndValues ...any) {
	log.Error(msg, keysAndValues...)
}

func Errorf(template string, args ...any) {
	log.Errorf(template, args...)
}

func Info(msg string, keysAndValues ...any) {
	log.Info(msg, keysAndValues...)
}

func Infof(template string, args ...any) {
	log.Infof(template, args...)
}

func Debug(msg string, keysAndValues ...any) {
	log.Debug(msg, keysAndValues...)
}

func Debugf(template string, args ...any) {
	log.Debugf(template, args...)
}
