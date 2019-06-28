package file

import "os"

type Option func(*config)

func WithFile(file *os.File) Option {
	return func(c *config) {
		c.file = file
	}
}

func WithErrorHandler(errorHandler ErrorHandler) Option {
	return func(c *config) {
		c.errorHandler = errorHandler
	}
}

type config struct {
	file         *os.File
	errorHandler ErrorHandler
}

func newConfig(opts ...Option) config {
	var c config
	defaultOpts := []Option{
		WithFile(os.Stdout),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}

	if c.errorHandler == nil {
		c.errorHandler = DefaultErrorHandler{
			file: c.file,
		}
	}

	return c
}
