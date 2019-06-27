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

func ParentSpanFromContext(ctx context.Context) (ParentSpan, bool) {
	span, ok := internal.SpanFromContext(ctx)
	if !ok {
		return ParentSpan{}, ok
	}

	return ParentSpan{
		ID:      span.ID(),
		TraceID: span.TraceID(),
	}, ok
}
