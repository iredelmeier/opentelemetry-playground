package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func ContextWithAttribute(ctx context.Context, key string, value string) context.Context {
	if _, ok := internal.AttributesFromContext(ctx); !ok {
		ctx = internal.ContextWithAttributes(ctx)
	}

	if attributes, ok := internal.AttributesFromContext(ctx); ok {
		attributes.Set(key, value)
	}

	return ctx
}

func AttributeFromContext(ctx context.Context, key string) (string, bool) {
	if attributes, ok := internal.AttributesFromContext(ctx); ok {
		return attributes.Get(key)
	}

	return "", false
}
