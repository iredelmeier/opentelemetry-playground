package trace

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/trace/internal"
)

const SpanIDSize = 8

type SpanID [SpanIDSize]byte

func SpanIDFromContext(ctx context.Context) (SpanID, bool) {
	span, ok := internal.SpanFromContext(ctx)
	if !ok {
		return SpanID{}, false
	}

	return span.ID(), true
}
