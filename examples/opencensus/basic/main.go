package main

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/bridges/opencensus"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	octrace "go.opencensus.io/trace"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	ocExporter := opencensus.NewExporter(opencensus.WithSpanExporter(exporter))

	octrace.RegisterExporter(ocExporter)
	octrace.ApplyConfig(octrace.Config{DefaultSampler: octrace.AlwaysSample()})

	ctx, parent := octrace.StartSpan(context.Background(), "parent")
	defer parent.End()

	_, child := octrace.StartSpan(ctx, "child")
	defer child.End()
}
