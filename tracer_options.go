package opentelemetry

type TracerOption func(*tracerConfig)

func WithSpanExporters(exporters ...SpanExporter) TracerOption {
	return func(c *tracerConfig) {
		c.exporters = exporters
	}
}

type tracerConfig struct {
	exporters []SpanExporter
}

func newTracerConfig(opts ...TracerOption) *tracerConfig {
	c := &tracerConfig{}
	var defaultOpts []TracerOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
