package log4go

import (
	"context"
	"sync"
)

var defaultLogger Logger

var loggers map[string]Logger
var loggerOnce sync.Once
var loggerMutex sync.Mutex

// SetLogger set logger
func SetLogger(name string, log Logger) {
	loggerOnce.Do(func() {
		loggers = make(map[string]Logger)
	})
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	loggers[name] = log
}

// GetLogger get Logger
func GetLogger(name string) Logger {
	return loggers[name]
}

// SetDefaultLogger 设置默认Logger
func SetDefaultLogger(dlog Logger) {
	defaultLogger = dlog
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Info(ctx, msg, fields...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Debug(ctx, msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Warn(ctx, msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Error(ctx, msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Panic(ctx, msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(ctx context.Context, msg string, fields ...Field) {
	if defaultLogger == nil {
		return
	}
	defaultLogger.Fatal(ctx, msg, fields...)
}

// Sync flushing any buffered log entries.
//
// Applications should take care to call Sync before exiting.
func Sync(ctx context.Context) {
	if defaultLogger != nil {
		defaultLogger.Sync(ctx)
	}
	for _, v := range loggers {
		v.Sync(ctx)
	}
}
