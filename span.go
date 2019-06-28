package opentelemetry

import (
	"context"
	"time"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Span struct {
	ID            SpanID
	TraceID       TraceID
	ParentID      SpanID
	OperationName string
	StartTime     time.Time
	Duration      time.Duration
	Tags          map[string]string
}

func StartSpan(ctx context.Context, operationName string, opts ...StartSpanOption) context.Context {
	c := newStartSpanConfig(opts...)

	spanOpts := []internal.StartSpanOption{
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(finishSpan),
	}

	if traceID, ok := TraceIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceID))
	}

	if parentID, ok := SpanIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithParentID(parentID))
	}

	spanOpts = append(spanOpts, c.opts...)

	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}

func FinishSpan(ctx context.Context) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		span.Finish(ctx)
	}
}

func finishSpan(ctx context.Context, span *internal.Span) {
	if exporter, ok := SpanExporterFromContext(ctx); ok {
		tags := make(map[string]string)

		if kv, ok := internal.KeyValuesFromContext(ctx); ok {
			for _, entry := range kv.Entries() {
				tags[entry.Key] = entry.Value
			}
		}

		exporter.ExportSpan(Span{
			ID:            span.ID(),
			TraceID:       span.TraceID(),
			ParentID:      span.ParentID(),
			OperationName: span.OperationName(),
			StartTime:     span.StartTime(),
			Duration:      time.Since(span.StartTime()),
			Tags:          tags,
		})
	}
}
