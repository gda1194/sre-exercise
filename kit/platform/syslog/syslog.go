package syslog

import (
	"context"
	"github.com/go-kit/log"
)

const keyLog = "logger"

type Logger struct {
	Logger log.Logger
}

// añadir el logger a gin.context
func Initialize(ctx context.Context, l log.Logger) context.Context {
	return context.WithValue(ctx, keyLog, Logger{Logger: l})
}

// obtener el logger del context.Context
func GetInstance(ctx context.Context) Logger {
	return ctx.Value(keyLog).(Logger)
}
