package pageSpeedInsight

import (
	"fmt"
	"google.golang.org/api/pagespeedonline/v5"
	"sre-exercise/internal"
	"sre-exercise/kit/platform/syslog"
)

type PerformanceIface interface {
	DoRequest(logger syslog.Logger, host string) (PerformanceMetrics, error)
}

type PageSpeedService struct {
	service *pagespeedonline.Service
}

func NewPageSpeedService(service *pagespeedonline.Service) PageSpeedService {
	return PageSpeedService{service: service}
}

type PerformanceMetrics struct {
	FirstContentfulPain    string
	SpeedIndex             string
	LargestContentFulPaint string
	TotalBlockingTime      string
	CumulativeLayoutShift  string
}

func (p PageSpeedService) DoRequest(logger syslog.Logger, host string) (PerformanceMetrics, error) {

	resp, err := p.service.Pagespeedapi.Runpagespeed(host).Do()
	if err != nil {
		logger.Error(err.Error())
		return PerformanceMetrics{}, err
	}

	performance := PerformanceMetrics{
		FirstContentfulPain:    formatMetric(resp.LighthouseResult.Audits[internal.FIRST_CONTENT_PAIN].NumericValue),
		SpeedIndex:             formatMetric(resp.LighthouseResult.Audits[internal.SPEED_INDEX].NumericValue),
		LargestContentFulPaint: formatMetric(resp.LighthouseResult.Audits[internal.LARGEST_CONTENFUL_PAINT].NumericValue),
		TotalBlockingTime:      formatMetric(resp.LighthouseResult.Audits[internal.TOTAL_BLOCKING_TIME].NumericValue),
		CumulativeLayoutShift:  formatMetric(resp.LighthouseResult.Audits[internal.CUMULATIVE_LAYOUT_SHIFT].NumericValue),
	}

	logger.Info(fmt.Sprintf("Performance to call %v -> %+v", host, performance))
	return performance, nil
}

func formatMetric(value float64) string {
	return fmt.Sprintf("%v ms", value)
}
