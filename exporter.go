package opentelemetry

type Exporter interface {
	Export(SpanData)
}
