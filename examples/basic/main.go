package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	ctx := trace.ContextWithSpanExporter(context.Background(), exporter)

	parentCtx := trace.StartSpan(ctx, "parent")
	defer trace.FinishSpan(parentCtx)

	childCtx := trace.StartSpan(parentCtx, "child")
	defer trace.FinishSpan(childCtx)
}
