package bootstrap

import (
	"context"
	"fmt"
	"github.com/dimiro1/health"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/robfig/cron/v3"
	otel "go.opentelemetry.io/otel"
	"google.golang.org/api/option"
	"google.golang.org/api/pagespeedonline/v5"
	"net/http"
	"os"
	"os/signal"
	"sre-exercise/internal"
	"sre-exercise/internal/openTelemetry"
	"sre-exercise/internal/pageSpeedInsight"
	"sre-exercise/internal/processor"
	"sre-exercise/kit/platform/server"
	"sre-exercise/kit/platform/syslog"
	"syscall"
)

func Run() error {

	// Init context
	logger := log.NewJSONLogger(os.Stderr)
	ctx := context.Background()

	// Init Service - PageSpeedInsights
	speedOnline, err := pagespeedonline.NewService(ctx, option.WithAPIKey(internal.PAGE_SPEED_API_KEY))
	if err != nil {
		return err
	}
	srvPerformance := pageSpeedInsight.NewPageSpeedService(speedOnline)

	// Init Instrumentations
	tracer, err := openTelemetry.GetTracer(ctx)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tracer)
	//tp := otel.Tracer(internal.TRACER_NAME)

	// Orchestra
	processorSrv := processor.NewProcessorService(srvPerformance, syslog.Logger{Logger: logger})

	// health
	errs := make(chan error, 2)
	srv := server.New("", "90", errs)
	handlerHealth := health.NewHandler()
	srv.RegisterRoute(http.MethodGet, "/health", gin.WrapH(handlerHealth))
	srv.Run()

	// cron
	scheduler := cron.New()
	scheduler.AddFunc("@every 1m", func() {
		// iniciamos la traza
		//_, span := tp.Start(ctx, internal.SPAN_NAME)

		//logger.Log(fmt.Sprintf("TraceID: %v", span.SpanContext().TraceID()))

		//srvPerformance.DoRequest(syslog.Logger{Logger: logger}, internal.MARCA_HOST)

		// terminamos la traza
		//span.End()
		processorSrv.WorkFlow(ctx)
		tracer.ForceFlush(ctx)
	})
	go func() {
		scheduler.Start()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		scheduler.Stop()
	}()
	logger = log.With(logger, "time", log.DefaultTimestamp)
	logger.Log("terminated", <-errs)
	tracer.Shutdown(ctx)

	return nil
}
