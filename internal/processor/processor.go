package processor

import (
	"sre-exercise/internal/pageSpeedInsight"
	"sre-exercise/kit/platform/syslog"
)

type ProcessorIface interface {
	WorkFlow() error
}

type ProcessorService struct {
	performanceSrv pageSpeedInsight.PerformanceIface
	logger         syslog.Logger
}

func NewProcessorService(performanceSrv pageSpeedInsight.PerformanceIface, logger syslog.Logger) *ProcessorService {
	return &ProcessorService{performanceSrv: performanceSrv, logger: logger}
}
