package opentelemetry

const (
	TraceIDSize = 16
	SpanIDSize  = 8
)

type SpanID [SpanIDSize]byte

type TraceID [TraceIDSize]byte

type Span struct {
	ID            SpanID
	TraceID       TraceID
	OperationName string
}
