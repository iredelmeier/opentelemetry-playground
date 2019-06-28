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
	Tags          map[string]string
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
		span.Finish(ctx)
	}
}

func finishSpan(ctx context.Context, span *internal.Span) {
	if exporter, ok := SpanExporterFromContext(ctx); ok {
		tags := make(map[string]string)

		if kv, ok := rootinternal.KeyValuesFromContext(ctx); ok {
			for _, entry := range kv.Entries() {
				tags[entry.Key] = entry.Value
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
			Tags:          tags,
		})
	}
}
