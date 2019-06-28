package trace

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/trace/internal"
)

const TraceIDSize = 16

type TraceID [TraceIDSize]byte

func TraceIDFromContext(ctx context.Context) (TraceID, bool) {
	span, ok := internal.SpanFromContext(ctx)
	if !ok {
		return TraceID{}, false
	}

	return span.TraceID(), true
}
