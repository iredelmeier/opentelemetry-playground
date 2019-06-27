package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Span struct {
	ID            SpanID
	TraceID       TraceID
	ParentID      SpanID
	OperationName string
	Tags          map[string]string
}

func newSpan(ctx context.Context, span *internal.Span) Span {
	tags := make(map[string]string)

	if kv, ok := internal.KeyValuesFromContext(ctx); ok {
		for _, entry := range kv.Entries() {
			tags[entry.Key] = entry.Value
		}
	}

	return Span{
		ID:            span.ID(),
		TraceID:       span.TraceID(),
		ParentID:      span.ParentID(),
		OperationName: span.OperationName(),
		Tags:          tags,
	}
}
