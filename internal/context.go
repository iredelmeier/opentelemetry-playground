package internal

import "context"

type spanCtxKey struct{}

func ContextWithSpan(ctx context.Context, span *Span) context.Context {
	return context.WithValue(ctx, spanCtxKey{}, span)
}

func SpanFromContext(ctx context.Context) *Span {
	span, ok := ctx.Value(spanCtxKey{}).(*Span)
	if !ok {
		return NewSpan()
	}

	return span
}
