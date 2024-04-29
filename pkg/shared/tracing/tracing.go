package tracing

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/errorreporting"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var errorClient *errorreporting.Client

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(ctx context.Context, projectId, service string) (tp *sdktrace.TracerProvider, er *errorreporting.Client) {
	exporter, err := texporter.New(texporter.WithProjectID(projectId))
	if err != nil {
		panic("texporter.New: " + err.Error())
	}

	// Identify your application using resource detection
	res, err := resource.New(ctx,
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(service),
		),
	)
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		panic(err.Error())
	} else if err != nil {
		panic("resource.New: " + err.Error())
	}

	tp = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	errorClient, err = errorreporting.NewClient(ctx, projectId, errorreporting.Config{
		ServiceName:    "errorreporting_quickstart",
		ServiceVersion: "0.0.0",
		OnError: func(err error) {
			log.Printf("Could not report the error: %v", err)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return tp, errorClient
}

func CreateSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.GetTracerProvider().Tracer("lutfi.dev")
	ctx, span := tracer.Start(ctx, name)

	return ctx, span
}

func PrintError(err error) {

	errorClient.Report(errorreporting.Entry{
		Error: err,
	})
	log.Print(err)
}
