package opentelemetry

type TraceContext struct {
	TraceID TraceID
	SpanID  SpanID
}
