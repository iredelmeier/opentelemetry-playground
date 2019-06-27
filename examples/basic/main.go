package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithSpanExporter(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)

	parentCtx := tracer.StartSpan(context.Background(), "parent")
	defer opentelemetry.FinishSpan(parentCtx)

	childCtx := tracer.StartSpan(parentCtx, "child")
	defer opentelemetry.FinishSpan(childCtx)
}
