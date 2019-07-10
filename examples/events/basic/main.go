package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/event"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	spanExporter := file.NewExporter()
	defer spanExporter.Close(context.Background())

	eventExporter := event.NewExporter(spanExporter)

	ctx := trace.ContextWithSpanEventExporter(context.Background(), eventExporter)

	parentCtx := trace.StartSpan(ctx, "parent")
	defer trace.FinishSpan(parentCtx)

	childCtx := trace.StartSpan(parentCtx, "child")
	defer trace.FinishSpan(childCtx)
}
