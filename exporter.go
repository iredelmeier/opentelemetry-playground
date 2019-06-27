package opentelemetry

type Exporter interface {
	Export(Span)
}
