package log4go

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {

	ctx := context.TODO()

	dlog := NewFileLogger(WithLevel("debug"))
	clog := NewConsoleLogger(WithLevel("debug"))

	glog := NewGroupLogger(dlog, clog)

	SetDefaultLogger(glog)

	defer Sync(ctx)

	Info(ctx, "test1")
	Debug(ctx, "debug test", Int("t2", 2))
	Panic(ctx, "panic")

}
