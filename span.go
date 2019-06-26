package opentelemetry

type finish func(*Span)

type Span struct {
	operationName string
	finish        finish
}

func (s *Span) Finish() {
	s.finish(s)
}
