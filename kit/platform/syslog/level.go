package syslog

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

const typeLog = "msg"
const typeSrv = "service"

// log.with deshabilitados debido a que en el bootstrap se hace una instancia directa

func (l Logger) Debug(params ...interface{}) {
	l.Logger = log.With(l.Logger, "time", log.DefaultTimestamp)
	level.Debug(l.Logger).Log(typeLog, params)
}
func (l Logger) Info(params ...interface{}) {
	l.Logger = log.With(l.Logger, "time", log.DefaultTimestamp)
	level.Info(l.Logger).Log(typeLog, params)
}
func (l Logger) Error(params ...interface{}) {
	l.Logger = log.With(l.Logger, "time", log.DefaultTimestamp)
	level.Error(l.Logger).Log(typeLog, params)
}
func (l Logger) Warn(params ...interface{}) {
	l.Logger = log.With(l.Logger, "time", log.DefaultTimestamp)
	level.Warn(l.Logger).Log(typeLog, params)
}
