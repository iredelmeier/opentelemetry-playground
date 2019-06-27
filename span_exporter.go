package opentelemetry

type SpanExporter interface {
	ExportSpan(Span)
}
