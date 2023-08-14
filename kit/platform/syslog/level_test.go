package syslog

import (
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ctxLogger_Debug(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	assert.NotNil(t, logger)

	ctl := Logger{Logger: logger}
	ctl.Debug("debug", "Test de log.debug")
}

func Test_ctxLogger_Error(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	assert.NotNil(t, logger)

	ctl := Logger{Logger: logger}
	ctl.Error("error", "Test de log.error")
}

func Test_ctxLogger_Info(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	assert.NotNil(t, logger)

	ctl := Logger{Logger: logger}
	ctl.Info("info", "Test de log.info")
}

func Test_ctxLogger_Warn(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	assert.NotNil(t, logger)

	ctl := Logger{Logger: logger}
	ctl.Warn("warn", "Test de log.warn")
}
