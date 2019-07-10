package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	exporterOpts := []lightstep.Option{
		lightstep.WithComponentName("basic"),
	}
	exporter, err := lightstep.NewExporter(exporterOpts...)
	if err != nil {
		panic(err)
	}
	defer exporter.Close(context.Background())

	ctx := trace.ContextWithSpanExporter(context.Background(), exporter)

	parentCtx := trace.StartSpan(ctx, "parent")
	defer trace.FinishSpan(parentCtx)

	childCtx := trace.StartSpan(parentCtx, "child")
	defer trace.FinishSpan(childCtx)
}
