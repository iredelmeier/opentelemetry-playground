package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/bridges/opencensus"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/trace"
	octrace "go.opencensus.io/trace"
)

func main() {
	exporter := trace.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	ocExporter := opencensus.NewExporter(opencensus.WithSpanExporter(exporter))

	octrace.RegisterExporter(ocExporter)
	octrace.ApplyConfig(octrace.Config{DefaultSampler: octrace.AlwaysSample()})

	ctx, parent := octrace.StartSpan(context.Background(), "parent")
	defer parent.End()

	_, child := octrace.StartSpan(ctx, "child")
	defer child.End()
}
