package trace

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/trace/internal"
)

const (
	SpanIDSize  = 8
	TraceIDSize = 16
)

type SpanID [SpanIDSize]byte

type TraceID [TraceIDSize]byte

type TraceContext struct {
	TraceID TraceID
	SpanID  SpanID
}

func TraceContextFromContext(ctx context.Context) (TraceContext, bool) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		return TraceContext{
			TraceID: span.TraceID(),
			SpanID:  span.ID(),
		}, true
	}

	return TraceContext{}, false
}
