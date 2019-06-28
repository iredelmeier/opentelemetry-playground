package trace

import (
	"context"
	"time"

	rootinternal "github.com/iredelmeier/opentelemetry-playground/internal"
	"github.com/iredelmeier/opentelemetry-playground/trace/internal"
)

type Span struct {
	TraceContext  TraceContext
	ParentID      SpanID
	OperationName string
	StartTime     time.Time
	FinishTime    time.Time
	Attributes    map[string]string
}

func StartSpan(ctx context.Context, operationName string, opts ...StartSpanOption) context.Context {
	c := newStartSpanConfig(opts...)

	spanOpts := []internal.StartSpanOption{
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(finishSpan),
	}

	if traceContext, ok := TraceContextFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceContext.TraceID))
		spanOpts = append(spanOpts, internal.WithParentID(traceContext.SpanID))
	}

	spanOpts = append(spanOpts, c.opts...)

	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}

func FinishSpan(ctx context.Context, opts ...FinishSpanOption) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		c := newFinishSpanConfig(opts...)

		span.Finish(ctx, c.opts...)
	}
}

func finishSpan(ctx context.Context, span internal.Span) {
	if exporter, ok := SpanExporterFromContext(ctx); ok {
		attributes := make(map[string]string)

		if a, ok := rootinternal.AttributesFromContext(ctx); ok {
			for _, attribute := range a.Entries() {
				attributes[attribute.Key] = attribute.Value
			}
		}

		exporter.ExportSpan(Span{
			TraceContext: TraceContext{
				TraceID: span.TraceID(),
				SpanID:  span.ID(),
			},
			ParentID:      span.ParentID(),
			OperationName: span.OperationName(),
			StartTime:     span.StartTime(),
			FinishTime:    span.FinishTime(),
			Attributes:    attributes,
		})
	}
}
