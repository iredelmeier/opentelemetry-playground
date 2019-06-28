package opentelemetry

import (
	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type StartSpanOption func(*startSpanConfig)

func WithID(id SpanID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.opts = append(c.opts, internal.WithID(id))
	}
}

func WithTraceID(traceID TraceID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.opts = append(c.opts, internal.WithTraceID(traceID))
	}
}

func WithParentID(parentID SpanID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.opts = append(c.opts, internal.WithParentID(parentID))
	}
}

type startSpanConfig struct {
	opts []internal.StartSpanOption
}

func newStartSpanConfig(opts ...StartSpanOption) *startSpanConfig {
	c := &startSpanConfig{}
	var defaultOpts []StartSpanOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
