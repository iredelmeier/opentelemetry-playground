package opentracing

import "github.com/iredelmeier/opentelemetry-playground/trace"

type TracerOption func(*tracerConfig)

func WithExporter(exporter trace.SpanExporter) TracerOption {
	return func(c *tracerConfig) {
		c.exporter = exporter
	}
}

type tracerConfig struct {
	exporter trace.SpanExporter
}

func newTracerConfig(opts ...TracerOption) tracerConfig {
	var c tracerConfig
	defaultOpts := []TracerOption{
		WithExporter(trace.NoopSpanExporter{}),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}

	return c
}
