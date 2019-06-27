package opentracing

import "github.com/iredelmeier/opentelemetry-playground"

type TracerOption func(*tracerConfig)

func WithExporter(exporter opentelemetry.SpanExporter) TracerOption {
	return func(c *tracerConfig) {
		c.exporter = exporter
	}
}

type tracerConfig struct {
	exporter opentelemetry.SpanExporter
}

func newTracerConfig(opts ...TracerOption) *tracerConfig {
	c := &tracerConfig{}
	defaultOpts := []TracerOption{
		WithExporter(opentelemetry.NoopSpanExporter{}),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
