package opentelemetry

type finish func(*Span)

type Span struct {
	id            SpanID
	traceID       TraceID
	operationName string
	finish        finish
}

func (s *Span) Finish() {
	s.finish(s)
}
