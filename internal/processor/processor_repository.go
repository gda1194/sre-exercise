package processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"sre-exercise/internal"
)

func (p ProcessorService) WorkFlow(ctx context.Context) error {

	// iniciamos la traza
	_, span := otel.Tracer(internal.TRACER_NAME).Start(ctx, internal.SPAN_NAME)

	p.logger.Info(fmt.Sprintf("TraceID: %v", span.SpanContext().TraceID()))

	_, err := p.performanceSrv.DoRequest(p.logger, internal.MARCA_HOST)
	if err != nil {
		return err
	}

	// terminamos la traza
	span.End()
	return nil

}
