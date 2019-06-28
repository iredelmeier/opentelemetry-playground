package opencensus

import "github.com/iredelmeier/opentelemetry-playground"

type Option func(*config)

func WithSpanExporter(spanExporter opentelemetry.SpanExporter) Option {
	return func(c *config) {
		c.spanExporter = spanExporter
	}
}

type config struct {
	spanExporter opentelemetry.SpanExporter
}

func newConfig(opts ...Option) *config {
	c := &config{}
	defaultOpts := []Option{
		WithSpanExporter(opentelemetry.NoopSpanExporter{}),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
