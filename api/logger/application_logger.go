package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

const (
	TransactionId = "transactionId"
)

var singletonLog Logger

func init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.MessageKey = "message"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := loggerConfig.Build(zap.WithCaller(true), zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal(err)
	}

	singletonLog = &zapLog{log: logger.Sugar()}
}

type zapLog struct {
	log *zap.SugaredLogger
}

func New() Logger {
	return singletonLog
}

func (zl *zapLog) Info(ctx context.Context, msg string, params ...interface{}) {
	l := zl.injectFields(ctx)
	if len(params) == 0 {
		l.Info(msg)
		return
	}
	l.Infof(msg, params...)
}

func (zl *zapLog) Debug(ctx context.Context, msg string, params ...interface{}) {
	l := zl.injectFields(ctx)
	if len(params) == 0 {
		l.Debug(msg)
		return
	}
	l.Debugf(msg, params...)
}

func (zl *zapLog) Warn(ctx context.Context, msg string, params ...interface{}) {
	l := zl.injectFields(ctx)
	if len(params) == 0 {
		l.Warn(msg)
		return
	}
	l.Warnf(msg, params...)
}

func (zl *zapLog) Error(ctx context.Context, msg string, params ...interface{}) {
	l := zl.injectFields(ctx)
	if len(params) == 0 {
		l.Error(msg)
		return
	}
	l.Errorf(msg, params...)
}

func (zl *zapLog) Fatal(ctx context.Context, msg string, params ...interface{}) {
	l := zl.injectFields(ctx)
	if len(params) == 0 {
		l.Fatal(msg)
		return
	}
	l.Fatalf(msg, params...)
}

func (zl *zapLog) injectFields(ctx context.Context) *zap.SugaredLogger {
	var tId string

	if ctx != nil {
		tId, _ = ctx.Value(TransactionId).(string)
	}
	return zl.log.WithOptions(zap.Fields(zap.String(TransactionId, tId)))
}
