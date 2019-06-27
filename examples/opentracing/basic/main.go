package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithSpanExporter(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)

	otTracer := bridge.NewTracer(bridge.WithOpenTelemetryTracer(tracer))

	parent := otTracer.StartSpan("parent")
	defer parent.Finish()

	child := otTracer.StartSpan("child", opentracing.ChildOf(parent.Context()))
	defer child.Finish()
}
