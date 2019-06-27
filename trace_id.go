package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

const TraceIDSize = 16

type TraceID [TraceIDSize]byte

func TraceIDFromContext(ctx context.Context) (TraceID, bool) {
	if state, ok := internal.StateFromContext(ctx); ok {
		return state.Span().TraceID(), true
	}

	return TraceID{}, false
}
