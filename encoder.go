package log4go

import (
	"encoding/json"
	"io"
	"time"

	"go.uber.org/zap/zapcore"
)

// ReflectedEncoder serializes log fields that can't be serialized with Zap's
// JSON encoder. These have the ReflectType field type.
// Use EncoderConfig.NewReflectedEncoder to set this.
type ReflectedEncoder interface {
	// Encode encodes and writes to the underlying data stream.
	Encode(interface{}) error
}

// PrimitiveArrayEncoder is the subset of the ArrayEncoder interface that deals
// only in Go's built-in types. It's included only so that Duration- and
// TimeEncoders cannot trigger infinite recursion.
type PrimitiveArrayEncoder zapcore.PrimitiveArrayEncoder

// A LevelEncoder serializes a Level to a primitive type.
type LevelEncoder func(Level, PrimitiveArrayEncoder)

// A TimeEncoder serializes a time.Time to a primitive type.
type TimeEncoder func(time.Time, PrimitiveArrayEncoder)

// A DurationEncoder serializes a time.Duration to a primitive type.
type DurationEncoder func(time.Duration, PrimitiveArrayEncoder)

type EntryCaller zapcore.EntryCaller

// A CallerEncoder serializes an EntryCaller to a primitive type.
type CallerEncoder func(EntryCaller, PrimitiveArrayEncoder)

// A NameEncoder serializes a period-separated logger name to a primitive
// type.
type NameEncoder func(string, PrimitiveArrayEncoder)

// LowercaseLevelEncoder serializes a Level to a lowercase string. For example,
// InfoLevel is serialized to "info".
func LowercaseLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
	enc.AppendString(zapcore.Level(l).String())
}

// TimeEncoderOfLayout returns TimeEncoder which serializes a time.Time using
// given layout.
func TimeEncoderOfLayout(layout string) TimeEncoder {
	return func(t time.Time, enc PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(layout))
	}
}

// StandTimeEncoder return TimeEncoder wich serializes a time.Time using layout 2006-01-02 15:04:05
func StandTimeEncoder() TimeEncoder {
	return TimeEncoderOfLayout("2006-01-02 15:04:05")
}

// MillisDurationEncoder serializes a time.Duration to an integer number of
// milliseconds elapsed.
func MillisDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Nanoseconds() / 1e6)
}

// FullCallerEncoder serializes a caller in /full/path/to/package/file:line
// format.
func FullCallerEncoder(caller EntryCaller, enc PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	enc.AppendString(zapcore.EntryCaller(caller).String())
}

// FullNameEncoder serializes the logger name as-is.
func FullNameEncoder(loggerName string, enc PrimitiveArrayEncoder) {
	enc.AppendString(loggerName)
}

// DefaultReflectedEncoder serializes the log in json
func DefaultReflectedEncoder(w io.Writer) ReflectedEncoder {
	enc := json.NewEncoder(w)
	// For consistency with our custom JSON encoder.
	enc.SetEscapeHTML(false)
	return enc
}
