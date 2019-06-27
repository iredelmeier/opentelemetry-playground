package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func FinishSpan(ctx context.Context) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		span.Finish(ctx)
	}
}
