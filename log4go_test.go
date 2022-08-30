package log4go

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {

	ctx := context.TODO()

	dlog := NewFileLogger(WithLevel("debug"), WithCaller(false), WithStack(false))
	clog := NewConsoleLogger(WithLevel("debug"), WithStack(false))

	glog := NewGroupLogger(dlog, clog)

	SetDefaultLogger(glog)

	// defer Sync(ctx)

	Info(ctx, "test1")
	Debug(ctx, "debug test", Int("t2", 2))
	Error(ctx, "error test", String("t", "www.baidu.com"))
	Warn(ctx, "warn test", String("w", "bilibili.com"))
	// Panic(ctx, "panic")

}

func BenchmarkConsoleLogger(b *testing.B) {
	ctx := context.TODO()

	dlog := NewFileLogger(WithFileName("/tmp/benchmark.log"), WithLevel("debug"), WithCaller(true), WithStack(true))
	// clog := NewConsoleLogger(WithLevel("debug"), WithStack(false))

	// glog := NewGroupLogger(dlog, clog)

	SetDefaultLogger(dlog)

	// defer Sync(ctx)

	for i := 0; i < b.N; i++ {
		Info(ctx, "test1")
		Debug(ctx, "debug test", Int("t2", 2), String("t", "www.baidu.com"), Int("t3", 2), String("t4", "www.baidu.com"), Int("t5", 2), String("t6", "www.baidu.com"))
		Error(ctx, "error test", String("t", "www.baidu.com"))
		Warn(ctx, "warn test", String("w", "bilibili.com"))
	}
}
