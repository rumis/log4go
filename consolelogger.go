package log4go

import (
	"context"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ConsoleLogger 基于zap的控制台日志
type ConsoleLogger struct {
	zap *zap.Logger
}

// NewConsoleLogger 创建默认的控制台日志
func NewConsoleLogger(oh ...OptionHandler) *ConsoleLogger {

	// 初始配置
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
	// 定义日志等级
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

	logger := zap.New(core, zapOpts...)

	return &ConsoleLogger{
		zap: logger,
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Info(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Info(msg)
		return
	}
	c.zap.Info(msg, FieldsConvert(fields)...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Debug(msg)
		return
	}
	c.zap.Debug(msg, FieldsConvert(fields)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Warn(msg)
		return
	}
	c.zap.Warn(msg, FieldsConvert(fields)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (c *ConsoleLogger) Error(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Error(msg)
		return
	}
	c.zap.Error(msg, FieldsConvert(fields)...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (c *ConsoleLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Panic(msg)
		return
	}
	c.zap.Panic(msg, FieldsConvert(fields)...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (c *ConsoleLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	if c.zap == nil {
		return
	}
	if len(fields) == 0 {
		c.zap.Fatal(msg)
		return
	}
	c.zap.Fatal(msg, FieldsConvert(fields)...)
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
