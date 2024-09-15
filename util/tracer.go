package util

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func newExporter(ctx context.Context) (exporter *jaeger.Exporter, err error) {
	exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(Configuration.Otel.Address)))
	return
}

func newTraceProvider(exp *jaeger.Exporter) *sdktrace.TracerProvider {
	resource, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewSchemaless(
				semconv.ServiceNameKey.String(fmt.Sprintf("%s/%s", Configuration.Meta.Application, Configuration.Meta.Version)),
				attribute.String("application", Configuration.Meta.Application),
				attribute.String("version", Configuration.Meta.Version),
				attribute.String("environment", Configuration.Meta.Environment),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)

}

func InboudGetSpan(requestContext *gin.Context, methodName string) (span_ctx context.Context, span trace.Span) {
	requestId := "1"
	span_ctx, span = Tracer.Start(requestContext.Request.Context(), methodName, trace.WithAttributes(
		attribute.String("requestId", requestId),
	))
	return
}

func init() {
	ctx := context.Background()

	exp, err := newExporter(ctx)
	if err != nil {
		panic("Failed Otel initialization")
	}

	tp := newTraceProvider(exp)

	otel.SetTracerProvider(tp)

	Tracer = tp.Tracer(fmt.Sprintf("%s/%s", Configuration.Meta.Application, Configuration.Meta.Version))
}
