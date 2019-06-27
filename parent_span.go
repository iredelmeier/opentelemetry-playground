package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type ParentSpan struct {
	ID      SpanID
	TraceID TraceID
}

func ContextWithParentSpan(ctx context.Context, parentSpan ParentSpan) context.Context {
	spanOpts := []internal.StartSpanOption{
		internal.WithID(parentSpan.ID),
		internal.WithTraceID(parentSpan.TraceID),
	}
	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}
