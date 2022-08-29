package log4go

import (
	"io"
	"strings"
)

type Options struct {
	MessageKey     string
	LevelKey       string
	TimeKey        string
	NameKey        string
	CallerKey      string
	FunctionKey    string
	StacktraceKey  string
	SkipLineEnding bool
	LineEnding     string
	// Configure the primitive representations of common complex types. For
	// example, some users may want all time.Times serialized as floating-point
	// seconds since epoch, while others may prefer ISO8601 strings.
	EncodeLevel    LevelEncoder
	EncodeTime     TimeEncoder
	EncodeDuration DurationEncoder
	EncodeCaller   CallerEncoder

	// Unlike the other primitive type encoders, EncodeName is optional. The
	// zero value falls back to FullNameEncoder.
	EncodeName NameEncoder

	// Configure the encoder for interface{} type objects.
	// If not provided, objects are encoded using json.Encoder
	NewReflectedEncoder func(io.Writer) ReflectedEncoder

	// Configures the field separator used by the console encoder. Defaults
	// to tab.
	ConsoleSeparator string

	// Level level
	Level Level
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool
}

type OptionHandler func(opt *Options)

func WithMessageKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.MessageKey = key
	}
}

func WithLevelKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.LevelKey = key
	}
}

func WithTimeKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.TimeKey = key
	}
}

func WithNameKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.NameKey = key
	}
}

func WithCallerKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.CallerKey = key
	}
}

func WithFunctionKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.FunctionKey = key
	}
}

func WithStacktraceKey(key string) OptionHandler {
	return func(opt *Options) {
		opt.StacktraceKey = key
	}
}

func WithSkipLineEnding(skip bool) OptionHandler {
	return func(opt *Options) {
		opt.SkipLineEnding = skip
	}
}

func WithLineEnding(end string) OptionHandler {
	return func(opt *Options) {
		opt.LineEnding = end
	}
}

func WithLevelEncoder(enc LevelEncoder) OptionHandler {
	return func(opt *Options) {
		opt.EncodeLevel = enc
	}
}

func WithDurationEncoder(enc DurationEncoder) OptionHandler {
	return func(opt *Options) {
		opt.EncodeDuration = enc
	}
}

func WithCallerEncoder(enc CallerEncoder) OptionHandler {
	return func(opt *Options) {
		opt.EncodeCaller = enc
	}
}

func WithTimeEncoder(enc TimeEncoder) OptionHandler {
	return func(opt *Options) {
		opt.EncodeTime = enc
	}
}
func WithNameEncoder(enc NameEncoder) OptionHandler {
	return func(opt *Options) {
		opt.EncodeName = enc
	}
}

func WithReflectedEncoder(fn func(io.Writer) ReflectedEncoder) OptionHandler {
	return func(opt *Options) {
		opt.NewReflectedEncoder = fn
	}
}

func WithConsoleSeparator(sep string) OptionHandler {
	return func(opt *Options) {
		opt.ConsoleSeparator = sep
	}
}

func WithFileName(filepath string) OptionHandler {
	return func(opt *Options) {
		opt.Filename = filepath
	}
}

func WithMaxAge(age int) OptionHandler {
	return func(opt *Options) {
		opt.MaxAge = age
	}
}

func WithMaxSize(size int) OptionHandler {
	return func(opt *Options) {
		opt.MaxSize = size
	}
}

func WithMaxBackups(maxb int) OptionHandler {
	return func(opt *Options) {
		opt.MaxBackups = maxb
	}
}

func WithLocalTime(lc bool) OptionHandler {
	return func(opt *Options) {
		opt.LocalTime = lc
	}
}

func WithCompress(compress bool) OptionHandler {
	return func(opt *Options) {
		opt.Compress = compress
	}
}

func WithLevel(level string) OptionHandler {
	l := strings.ToLower(level)
	return func(opt *Options) {
		switch l {
		case "debug":
			opt.Level = DebugLevel
		case "info":
			opt.Level = InfoLevel
		case "warn":
			opt.Level = WarnLevel
		case "error":
			opt.Level = ErrorLevel
		case "panic":
			opt.Level = PanicLevel
		case "fatal":
			opt.Level = FatalLevel
		default: // default info
			opt.Level = InfoLevel
		}
	}
}

// DefaultOption 默认配置
func DefaultOption() Options {
	return Options{
		Filename:            "/tmp/lumberjack.log",
		MaxSize:             100,
		MaxAge:              0,
		MaxBackups:          0,
		LocalTime:           false,
		Compress:            false,
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "name",
		CallerKey:           "caller",
		FunctionKey:         "func",
		StacktraceKey:       "stack",
		SkipLineEnding:      false,
		LineEnding:          "\n",
		EncodeLevel:         LowercaseLevelEncoder,
		EncodeTime:          StandTimeEncoder(),
		EncodeDuration:      MillisDurationEncoder,
		EncodeCaller:        FullCallerEncoder,
		EncodeName:          FullNameEncoder,
		ConsoleSeparator:    "\t",
		NewReflectedEncoder: DefaultReflectedEncoder,
	}
}
