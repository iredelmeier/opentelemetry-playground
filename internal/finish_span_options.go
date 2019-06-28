package internal

import "time"

type FinishSpanOption func(*finishSpanConfig)

func WithFinishTime(finishTime time.Time) FinishSpanOption {
	return func(c *finishSpanConfig) {
		c.finishTime = finishTime
	}
}

type finishSpanConfig struct {
	finishTime time.Time
}

func newFinishSpanConfig(opts ...FinishSpanOption) *finishSpanConfig {
	c := &finishSpanConfig{}
	var defaultOpts []FinishSpanOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	if c.finishTime.IsZero() {
		c.finishTime = time.Now()
	}

	return c
}
