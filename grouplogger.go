package log4go

import "context"

// GroupLogger multi logger
type GroupLogger struct {
	loggers []Logger
}

// NewGroupLogger new multi logger
func NewGroupLogger(logger ...Logger) *GroupLogger {
	return &GroupLogger{
		loggers: logger,
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (g *GroupLogger) Info(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Info(ctx, msg, fields...)
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (g *GroupLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Debug(ctx, msg, fields...)
	}
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (g *GroupLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Warn(ctx, msg, fields...)
	}
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (g *GroupLogger) Error(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Error(ctx, msg, fields...)
	}
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (g *GroupLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Panic(ctx, msg, fields...)
	}
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (g *GroupLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	for _, l := range g.loggers {
		l.Fatal(ctx, msg, fields...)
	}
}

// Sync flushing any buffered log entries.
//
// Applications should take care to call Sync before exiting.
func (g *GroupLogger) Sync(ctx context.Context) {
	for _, l := range g.loggers {
		l.Sync(ctx)
	}
}
