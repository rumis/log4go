# log4go

logging in Go, based on [zap](https://github.com/uber-go/zap).

## Installation

`go get github.com/rumis/log4go`

Note that log4go only supports the two most recent minor versions of Go, Because of zap.

## Quick Start

```go
ctx := context.TODO()
consoleLogger := log4go.NewConsoleLogger()
log4go.SetDefaultLogger(consoleLogger)
defer log4go.Sync(ctx)
log4go.Info(ctx, "failed to fetch URL",
    log4go.String("url",url),
    log4go.Int("attempt", 3),
    log4go.Duration("backoff", time.Second))
```

See the [documentation](https://pkg.go.dev/github.com/rumis/log4go) for more details.


<hr>

Released under the [MIT License](LICENSE.txt).