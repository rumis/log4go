package log4go

import (
	"context"
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// FileLogger 基于zap的文件日志
type FileLogger struct {
	zap *zap.Logger
}

// NewFileLogger 创建默认的文件日志
func NewFileLogger(oh ...OptionHandler) *FileLogger {

	// 初始配置
	opts := DefaultOption()
	for _, fn := range oh {
		fn(&opts)
	}
	// 定义文件写入&拆分逻辑
	hook := lumberjack.Logger{
		Filename:   opts.Filename,
		MaxSize:    opts.MaxSize,
		MaxAge:     opts.MaxAge,
		MaxBackups: opts.MaxBackups,
		Compress:   opts.Compress,
		LocalTime:  opts.LocalTime,
	}
	write := zapcore.AddSync(&hook)

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
		zapcore.NewJSONEncoder(encoderConfig),
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

	return &FileLogger{
		zap: logger,
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (f *FileLogger) Info(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Info(msg)
		return
	}
	f.zap.Info(msg, FieldsConvert(fields)...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (f *FileLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Debug(msg)
		return
	}
	f.zap.Debug(msg, FieldsConvert(fields)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (f *FileLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Warn(msg)
		return
	}
	f.zap.Warn(msg, FieldsConvert(fields)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (f *FileLogger) Error(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Error(msg)
		return
	}
	f.zap.Error(msg, FieldsConvert(fields)...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (f *FileLogger) Panic(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Panic(msg)
		return
	}
	f.zap.Panic(msg, FieldsConvert(fields)...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (f *FileLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	if f.zap == nil {
		return
	}
	if len(fields) == 0 {
		f.zap.Fatal(msg)
		return
	}
	f.zap.Fatal(msg, FieldsConvert(fields)...)
}

// Sync flushing any buffered log entries.
//
// Applications should take care to call Sync before exiting.
func (f *FileLogger) Sync(ctx context.Context) {
	if f.zap == nil {
		return
	}
	f.zap.Sync()
}
