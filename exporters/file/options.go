package file

import "os"

type Option func(*Config)

func WithFile(file *os.File) Option {
	return func(c *Config) {
		c.file = file
	}
}

func WithErrorHandler(errorHandler ErrorHandler) Option {
	return func(c *Config) {
		c.errorHandler = errorHandler
	}
}

type Config struct {
	file         *os.File
	errorHandler ErrorHandler
}

func newConfig(opts ...Option) *Config {
	c := &Config{}
	defaultOpts := []Option{
		WithFile(os.Stdout),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	if c.errorHandler == nil {
		c.errorHandler = &DefaultErrorHandler{
			file: c.file,
		}
	}

	return c
}
