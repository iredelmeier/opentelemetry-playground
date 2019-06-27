package opentelemetry

type SpanExporter interface {
	ExportSpan(Span)
}

type NoopSpanExporter struct{}

func (NoopSpanExporter) ExportSpan(Span) {}
