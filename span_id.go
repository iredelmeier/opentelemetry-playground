package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

const SpanIDSize = 8

type SpanID [SpanIDSize]byte

func SpanIDFromContext(ctx context.Context) (SpanID, bool) {
	if state, ok := internal.StateFromContext(ctx); ok {
		return state.Span().ID(), true
	}

	return SpanID{}, false
}
