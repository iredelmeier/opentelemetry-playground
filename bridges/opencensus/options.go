package opencensus

import "github.com/iredelmeier/opentelemetry-playground/trace"

type Option func(*config)

func WithSpanExporter(spanExporter trace.SpanExporter) Option {
	return func(c *config) {
		c.spanExporter = spanExporter
	}
}

type config struct {
	spanExporter trace.SpanExporter
}

func newConfig(opts ...Option) config {
	var c config
	defaultOpts := []Option{
		WithSpanExporter(trace.NoopSpanExporter{}),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}

	return c
}
