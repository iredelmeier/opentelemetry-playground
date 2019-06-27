package opentelemetry

type TracerOption func(*tracerConfig)

func WithSpanExporter(exporter SpanExporter) TracerOption {
	return func(c *tracerConfig) {
		c.exporter = exporter
	}
}

type tracerConfig struct {
	exporter SpanExporter
}

func newTracerConfig(opts ...TracerOption) *tracerConfig {
	c := &tracerConfig{}
	defaultOpts := []TracerOption{
		WithSpanExporter(NoopSpanExporter{}),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
