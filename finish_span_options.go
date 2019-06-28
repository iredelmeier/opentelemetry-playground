package opentelemetry

import (
	"time"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type FinishSpanOption func(*finishSpanConfig)

func WithFinishTime(finishTime time.Time) FinishSpanOption {
	return func(c *finishSpanConfig) {
		c.opts = append(c.opts, internal.WithFinishTime(finishTime))
	}
}

type finishSpanConfig struct {
	opts []internal.FinishSpanOption
}

func newFinishSpanConfig(opts ...FinishSpanOption) *finishSpanConfig {
	c := &finishSpanConfig{}
	var defaultOpts []FinishSpanOption

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}
