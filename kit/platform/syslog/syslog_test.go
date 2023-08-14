package syslog

import (
	context2 "context"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	ctx := context2.TODO()

	ctLog := Initialize(ctx, logger)
	testLog := ctLog.Value("logger").(Logger)
	assert.NotNil(t, testLog)
}

func TestGetInstance(t *testing.T) {
	logger := log.NewJSONLogger(os.Stderr)
	ctx := context2.TODO()

	ctLog := Initialize(ctx, logger)
	l := GetInstance(ctLog)
	assert.Equal(t, Logger{Logger: logger}, l)
}
