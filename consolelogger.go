package log4go

import (
	"context"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ConsoleLogger console logger base on zap
type ConsoleLogger struct {
	zap       *zap.Logger
	extfields []Field
}

// NewConsoleLogger create a new ConsoleLogger
func NewConsoleLogger(oh ...OptionHandler) *ConsoleLogger {

	// initialize config
	opts := DefaultOption()
	for _, fn := range oh {
		fn(&opts)
	}

	write := zapcore.AddSync(os.Stdout)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        opts.TimeKey,
		LevelKey:       opts.LevelKey,
		NameKey:        opts.NameKey,
		CallerKey:      opts.CallerKey,
		MessageKey:     opts.MessageKey,
		StacktraceKey:  opts.StacktraceKey,
		SkipLineEnding: opts.SkipLineEnding,
		LineEnding:     opts.LineEnding,
		FunctionKey:    opts.FunctionKey,
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			opts.EncodeLevel(l, enc)
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			opts.EncodeTime(t, enc)
		},
		EncodeDuration: func(td time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			opts.EncodeDuration(td, enc)
		},
		EncodeCaller: func(ec zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			opts.EncodeCaller(EntryCaller(ec), enc)
		},
		EncodeName: func(name string, enc zapcore.PrimitiveArrayEncoder) {
			opts.EncodeName(name, enc)
		},
		NewReflectedEncoder: func(w io.Writer) zapcore.ReflectedEncoder {
			return opts.NewReflectedEncoder(w)
		},
		ConsoleSeparator: opts.ConsoleSeparator,
	}
	// log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(opts.Level)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		write,
		atomicLevel)

	zapOpts := make([]zap.Option, 0)
	if opts.WithCaller {
		zapOpts = append(zapOpts, zap.AddCaller())
	}
	if opts.WithStack {
		zapOpts = append(zapOpts, zap.AddStacktrace(DebugLevel))
	}
	// zap core
	logger := zap.New(core, zapOpts...)

	return &ConsoleLogger{
		zap:       logger,
		extfields: opts.ExtFields,
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Info(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, InfoLevel, msg, fields...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, DebugLevel, msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, WarnLevel, msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Error(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, ErrorLevel, msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (c *ConsoleLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, PanicLevel, msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (c *ConsoleLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	c.Log(ctx, FatalLevel, msg, fields...)
}

// Log logs a message at the specified level. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Log(ctx context.Context, lvl Level, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	// static extend fields
	fields = append(fields, c.extfields...)

	// context extend fields
	cval := ctx.Value(ContextFieldsKey)
	if cval != nil {
		if cfields, ok := cval.([]Field); ok {
			fields = append(fields, cfields...)
		}
	}
	// write
	if ce := c.zap.Check(lvl, msg); ce != nil {
		if len(fields) == 0 {
			ce.Write()
			return
		}
		ce.Write(FieldsConvert(fields)...)
	}
}

// Sync flushing any buffered log entries.
//
// Applications should take care to call Sync before exiting.
func (c *ConsoleLogger) Sync(ctx context.Context) {
	if c.zap == nil {
		return
	}
	c.zap.Sync()
}
