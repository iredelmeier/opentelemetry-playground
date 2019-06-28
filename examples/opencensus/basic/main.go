package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/bridges/opencensus"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"go.opencensus.io/trace"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	ocExporter := opencensus.NewExporter(opencensus.WithSpanExporter(exporter))

	trace.RegisterExporter(ocExporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	ctx, parent := trace.StartSpan(context.Background(), "parent")
	defer parent.End()

	_, child := trace.StartSpan(ctx, "child")
	defer child.End()
}
