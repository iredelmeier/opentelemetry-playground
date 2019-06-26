package opentelemetry

type TracerOption func(*tracerConfig)

func WithExporters(exporters ...Exporter) TracerOption {
	return func(c *tracerConfig) {
		c.exporters = exporters
	}
}

type tracerConfig struct {
	exporters []Exporter
}

func newTracerConfig(opts ...TracerOption) *tracerConfig {
	c := &tracerConfig{}
	var defaultOpts []TracerOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
