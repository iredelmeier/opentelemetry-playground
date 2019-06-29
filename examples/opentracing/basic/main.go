package main

import (
	"context"

	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	otTracer := bridge.NewTracer(bridge.WithExporter(exporter))

	parent := otTracer.StartSpan("parent")
	defer parent.Finish()

	child := otTracer.StartSpan("child", opentracing.ChildOf(parent.Context()))
	defer child.Finish()
}
