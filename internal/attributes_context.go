package internal

import (
	"context"
)

type attributesCtxKey struct{}

func ContextWithAttributes(ctx context.Context) context.Context {
	return context.WithValue(ctx, attributesCtxKey{}, NewAttributes())
}

func AttributesFromContext(ctx context.Context) (Attributes, bool) {
	attributes, ok := ctx.Value(attributesCtxKey{}).(Attributes)

	return attributes, ok
}
