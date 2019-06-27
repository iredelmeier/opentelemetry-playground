package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func FinishSpan(ctx context.Context) {
	if state, ok := internal.StateFromContext(ctx); ok {
		state.Span().Finish()
	}
}
