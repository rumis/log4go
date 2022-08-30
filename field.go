package log4go

import (
	"math"
	"time"

	"go.uber.org/zap/zapcore"
)

type Field zapcore.Field
type Fields []zapcore.Field

type ContextField struct{}

// ContextFieldsKey key of context extend fields
var ContextFieldsKey = ContextField{}

var (
	_minTimeInt64 = time.Unix(0, math.MinInt64)
	_maxTimeInt64 = time.Unix(0, math.MaxInt64)
)

const (
	// UnknownType is the default field type. Attempting to add it to an encoder will panic.
	UnknownType = zapcore.UnknownType
	// BoolType indicates that the field carries a bool.
	BoolType = zapcore.BoolType
	// ByteStringType indicates that the field carries UTF-8 encoded bytes.
	ByteStringType = zapcore.ByteStringType
	// DurationType indicates that the field carries a time.Duration.
	DurationType = zapcore.DurationType
	// Float64Type indicates that the field carries a float64.
	Float64Type = zapcore.Float64Type
	// Int64Type indicates that the field carries an int64.
	Int64Type = zapcore.Int64Type
	// StringType indicates that the field carries a string.
	StringType = zapcore.StringType
	TimeType   = zapcore.TimeType
	// TimeFullType indicates that the field carries a time.Time stored as-is.
	TimeFullType = zapcore.TimeFullType
	// Uint64Type indicates that the field carries a uint64.
	Uint64Type = zapcore.Uint64Type
	// ErrorType indicates that the field carries an error.
	ErrorType = zapcore.ErrorType
)

// FieldsConvert convert Field to zapcore.Field
func FieldsConvert(fields []Field) []zapcore.Field {
	f := make([]zapcore.Field, 0, len(fields))
	for _, v := range fields {
		f = append(f, zapcore.Field(v))
	}
	return f
}

// Bool constructs a field that carries a bool.
func Bool(key string, val bool) Field {
	var ival int64
	if val {
		ival = 1
	}
	return Field{Key: key, Type: BoolType, Integer: ival}
}

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
// Binary.
func ByteString(key string, val []byte) Field {
	return Field{Key: key, Type: ByteStringType, Interface: val}
}

// Float64 constructs a field that carries a float64. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
func Float64(key string, val float64) Field {
	return Field{Key: key, Type: Float64Type, Integer: int64(math.Float64bits(val))}
}

// Int constructs a field with the given key and value.
func Int(key string, val int) Field {
	return Int64(key, int64(val))
}

// Int64 constructs a field with the given key and value.
func Int64(key string, val int64) Field {
	return Field{Key: key, Type: Int64Type, Integer: val}
}

// String constructs a field with the given key and value.
func String(key string, val string) Field {
	return Field{Key: key, Type: StringType, String: val}
}

// Uint constructs a field with the given key and value.
func Uint(key string, val uint) Field {
	return Uint64(key, uint64(val))
}

// Uint64 constructs a field with the given key and value.
func Uint64(key string, val uint64) Field {
	return Field{Key: key, Type: Uint64Type, Integer: int64(val)}
}

// Time constructs a Field with the given key and value. The encoder
// controls how the time is serialized.
func Time(key string, val time.Time) Field {
	if val.Before(_minTimeInt64) || val.After(_maxTimeInt64) {
		return Field{Key: key, Type: TimeFullType, Interface: val}
	}
	return Field{Key: key, Type: TimeType, Integer: val.UnixNano(), Interface: val.Location()}
}

// Duration constructs a field with the given key and value. The encoder
// controls how the duration is serialized.
func Duration(key string, val time.Duration) Field {
	return Field{Key: key, Type: DurationType, Integer: int64(val)}
}
