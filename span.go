package opentelemetry

import "github.com/iredelmeier/opentelemetry-playground/internal"

const (
	TraceIDSize = 16
	SpanIDSize  = 8
)

type SpanID [SpanIDSize]byte

type TraceID [TraceIDSize]byte

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
