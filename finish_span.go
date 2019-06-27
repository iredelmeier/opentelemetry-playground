package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func FinishSpan(ctx context.Context) {
	span, ok := internal.SpanFromContext(ctx)
	if ok {
		span.Finish()
	}
}
