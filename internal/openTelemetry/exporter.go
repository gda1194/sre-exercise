package openTelemetry

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"sre-exercise/internal"
)

func GetTracer(ctx context.Context) (*trace.TracerProvider, error) {

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(internal.GRAFANA_TEMPO_ENDPOINT),
		otlptracegrpc.WithHeaders(
			map[string]string{
				"Content-Type":  "application/grpc",
				"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(internal.GRAFANA_TEMPO_API_KEY)),
			},
		),
		otlptracegrpc.WithTLSCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
	)
	if err != nil {
		return nil, err
	}

	return NewTracerService(exporter).NewTrace(), nil
}
