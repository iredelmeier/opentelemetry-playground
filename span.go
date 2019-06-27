package opentelemetry

import "github.com/iredelmeier/opentelemetry-playground/internal"

type Span struct {
	ID            SpanID
	TraceID       TraceID
	ParentID      SpanID
	OperationName string
}

func newSpan(span *internal.Span) Span {
	return Span{
		ID:            span.ID(),
		TraceID:       span.TraceID(),
		ParentID:      span.ParentID(),
		OperationName: span.OperationName(),
	}
}
