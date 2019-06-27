package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	ctx := opentelemetry.ContextWithSpanExporter(context.Background(), exporter)

	parentCtx := opentelemetry.StartSpan(ctx, "parent")
	defer opentelemetry.FinishSpan(parentCtx)

	childCtx := opentelemetry.StartSpan(parentCtx, "child")
	defer opentelemetry.FinishSpan(childCtx)
}
