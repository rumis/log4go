package log4go

import "context"

// Logger logger
type Logger interface {
	// Info logs a message at InfoLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	Info(ctx context.Context, msg string, fields ...Field)

	// Debug logs a message at DebugLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	Debug(ctx context.Context, msg string, fields ...Field)

	// Warn logs a message at WarnLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	Warn(ctx context.Context, msg string, fields ...Field)

	// Error logs a message at ErrorLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	Error(ctx context.Context, msg string, fields ...Field)

	// Panic logs a message at PanicLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	//
	// The logger then panics, even if logging at PanicLevel is disabled.
	Panic(ctx context.Context, msg string, fields ...Field)

	// Fatal logs a message at FatalLevel. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	//
	// The logger then calls os.Exit(1), even if logging at FatalLevel is
	// disabled.
	Fatal(ctx context.Context, msg string, fields ...Field)

	// Sync flushing any buffered log entries.
	//
	// Applications should take care to call Sync before exiting.
	Sync(ctx context.Context)
}
