package internal

import "context"

type ctxKey struct{}

func ContextWithState(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, NewState())
}

func StateFromContext(ctx context.Context) (*State, bool) {
	state, ok := ctx.Value(ctxKey{}).(*State)

	return state, ok
}
