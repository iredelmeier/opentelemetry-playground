package internal

import (
	"context"
	"sync"
)

type kvCtxKey struct{}

func ContextWithKeyValues(ctx context.Context) context.Context {
	return context.WithValue(ctx, kvCtxKey{}, &KeyValues{
		lock: &sync.RWMutex{},
		kv:   make(map[string]string),
	})
}

func KeyValuesFromContext(ctx context.Context) (*KeyValues, bool) {
	kv, ok := ctx.Value(kvCtxKey{}).(*KeyValues)

	return kv, ok
}
