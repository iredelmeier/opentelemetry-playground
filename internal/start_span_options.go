package internal

type StartSpanOption func(*startSpanConfig)

func WithID(id [8]byte) StartSpanOption {
	return func(c *startSpanConfig) {
		c.id = id
	}
}

func WithTraceID(traceID [16]byte) StartSpanOption {
	return func(c *startSpanConfig) {
		c.traceID = traceID
	}
}

func WithParentID(parentID [8]byte) StartSpanOption {
	return func(c *startSpanConfig) {
		c.parentID = parentID
	}
}

func WithOperationName(operationName string) StartSpanOption {
	return func(c *startSpanConfig) {
		c.operationName = operationName
	}
}

func WithFinishSpan(finishSpan FinishSpan) StartSpanOption {
	return func(c *startSpanConfig) {
		c.finishSpan = finishSpan
	}
}

type startSpanConfig struct {
	id            [8]byte
	traceID       [16]byte
	parentID      [8]byte
	operationName string
	finishSpan    FinishSpan
}

func newStartSpanConfig(opts ...StartSpanOption) *startSpanConfig {
	c := &startSpanConfig{}
	defaultOpts := []StartSpanOption{
		WithFinishSpan(defaultFinishSpan),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
