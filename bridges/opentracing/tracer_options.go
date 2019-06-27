package opentracing

import "github.com/iredelmeier/opentelemetry-playground"

type TracerOption func(*tracerConfig)

func WithOpenTelemetryTracer(openTelemetryTracer *opentelemetry.Tracer) TracerOption {
	return func(c *tracerConfig) {
		c.openTelemetryTracer = openTelemetryTracer
	}
}

type tracerConfig struct {
	openTelemetryTracer *opentelemetry.Tracer
}

func newTracerConfig(opts ...TracerOption) *tracerConfig {
	c := &tracerConfig{}
	defaultOpts := []TracerOption{
		WithOpenTelemetryTracer(opentelemetry.NewTracer()),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
