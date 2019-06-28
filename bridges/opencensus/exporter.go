package opencensus

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/trace"
	octrace "go.opencensus.io/trace"
)

type Exporter struct {
	spanExporter trace.SpanExporter
}

func NewExporter(opts ...Option) Exporter {
	c := newConfig(opts...)

	return Exporter{
		spanExporter: c.spanExporter,
	}
}

func (e Exporter) ExportSpan(span *octrace.SpanData) {
	startOpts := []trace.StartSpanOption{
		trace.WithID(trace.SpanID(span.SpanContext.SpanID)),
		trace.WithTraceID(trace.TraceID(span.SpanContext.TraceID)),
		trace.WithParentID(trace.SpanID(span.ParentSpanID)),
		trace.WithStartTime(span.StartTime),
	}
	ctx := trace.ContextWithSpanExporter(context.Background(), e.spanExporter)

	ctx = trace.StartSpan(ctx, span.Name, startOpts...)

	finishOpts := []trace.FinishSpanOption{
		trace.WithFinishTime(span.EndTime),
	}

	trace.FinishSpan(ctx, finishOpts...)
}
