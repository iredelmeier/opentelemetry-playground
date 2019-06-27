package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
)

func main() {
	exporter := file.NewExporter()
	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithExporters(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)
	defer tracer.Close(context.Background())

	parentCtx := tracer.StartSpan(context.Background(), "parent")
	defer opentelemetry.FinishSpan(parentCtx)

	childCtx := tracer.StartSpan(parentCtx, "child")
	defer opentelemetry.FinishSpan(childCtx)
}
