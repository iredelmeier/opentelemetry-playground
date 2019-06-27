package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func FinishSpan(ctx context.Context) {
	span := internal.SpanFromContext(ctx)

	span.Finish()
}
