package opencensus

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"go.opencensus.io/trace"
)

type Exporter struct {
	spanExporter opentelemetry.SpanExporter
}

func NewExporter(opts ...Option) *Exporter {
	c := newConfig(opts...)

	return &Exporter{
		spanExporter: c.spanExporter,
	}
}

func (e *Exporter) ExportSpan(span *trace.SpanData) {
	startOpts := []opentelemetry.StartSpanOption{
		opentelemetry.WithID(opentelemetry.SpanID(span.SpanContext.SpanID)),
		opentelemetry.WithTraceID(opentelemetry.TraceID(span.SpanContext.TraceID)),
		opentelemetry.WithParentID(opentelemetry.SpanID(span.ParentSpanID)),
		opentelemetry.WithStartTime(span.StartTime),
	}
	ctx := opentelemetry.ContextWithSpanExporter(context.Background(), e.spanExporter)

	ctx = opentelemetry.StartSpan(ctx, span.Name, startOpts...)

	finishOpts := []opentelemetry.FinishSpanOption{
		opentelemetry.WithFinishTime(span.EndTime),
	}

	opentelemetry.FinishSpan(ctx, finishOpts...)
}
