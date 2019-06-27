package opentelemetry

type StartSpanOption func(*startSpanConfig)

func WithID(id SpanID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.id = id
	}
}

func WithTraceID(traceID TraceID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.traceID = traceID
	}
}

func WithParentID(parentID SpanID) StartSpanOption {
	return func(c *startSpanConfig) {
		c.parentID = parentID
	}
}

type startSpanConfig struct {
	id       SpanID
	traceID  TraceID
	parentID SpanID
}

func newStartSpanConfig(opts ...StartSpanOption) *startSpanConfig {
	c := &startSpanConfig{}
	var defaultOpts []StartSpanOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
