package opentelemetry

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

func ContextWithKeyValue(ctx context.Context, key string, value string) context.Context {
	if _, ok := internal.KeyValuesFromContext(ctx); !ok {
		ctx = internal.ContextWithKeyValues(ctx)
	}

	if kv, ok := internal.KeyValuesFromContext(ctx); ok {
		kv.Set(key, value)
	}

	return ctx
}

func KeyValueFromContext(ctx context.Context, key string) (string, bool) {
	if kv, ok := internal.KeyValuesFromContext(ctx); ok {
		return kv.Get(key)
	}

	return "", false
}
