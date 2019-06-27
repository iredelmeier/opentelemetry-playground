package opentelemetry

import (
	"context"
	"time"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Span struct {
	ID            SpanID
	TraceID       TraceID
	ParentID      SpanID
	OperationName string
	StartTime     time.Time
	Duration      time.Duration
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
		StartTime:     span.StartTime(),
		Duration:      time.Since(span.StartTime()),
		Tags:          tags,
	}
}
