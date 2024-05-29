package logger

import (
	"context"
)

type Logger interface {
	Info(ctx context.Context, msg string, params ...interface{})
	Debug(ctx context.Context, msg string, params ...interface{})
	Error(ctx context.Context, msg string, params ...interface{})
	Warn(ctx context.Context, msg string, params ...interface{})
	Fatal(ctx context.Context, msg string, params ...interface{})
}
