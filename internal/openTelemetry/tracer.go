package openTelemetry

import (
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"sre-exercise/internal"
)

type TraceIface interface {
	NewTrace() *trace.TracerProvider
}

type TracerService struct {
	exporter trace.SpanExporter
}

func NewTracerService(exporter trace.SpanExporter) *TracerService {
	return &TracerService{exporter: exporter}
}

func (ts *TracerService) NewTrace() *trace.TracerProvider {
	tp := trace.NewTracerProvider(
		trace.WithBatcher(ts.exporter),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(internal.GRAFANA_TEMPO_SERVICE_NAME),
				semconv.ServiceVersion(internal.GRAFANA_TEMPO_SERVICE_VERSION),
			),
		),
	)
	return tp
}
