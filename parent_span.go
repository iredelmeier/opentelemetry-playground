package opentelemetry

type ParentSpan struct {
	ID      SpanID
	TraceID TraceID
}
