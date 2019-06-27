package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Tracer struct {
	exporter SpanExporter
}

func NewTracer(opts ...TracerOption) *Tracer {
	c := newTracerConfig(opts...)

	return &Tracer{
		exporter: c.exporter,
	}
}

func (t *Tracer) StartSpan(ctx context.Context, operationName string, opts ...StartSpanOption) context.Context {
	c := newStartSpanConfig(opts...)

	spanOpts := []internal.StartSpanOption{
		internal.WithID(c.id),
		internal.WithTraceID(c.traceID),
		internal.WithParentID(c.parentID),
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(t.finishSpan),
	}

	if traceID, ok := TraceIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceID))
	}

	if parentID, ok := SpanIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithParentID(parentID))
	}

	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}

func (t *Tracer) finishSpan(ctx context.Context, span *internal.Span) {
	t.exporter.ExportSpan(newSpan(ctx, span))
}
