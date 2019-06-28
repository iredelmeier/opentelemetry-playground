package trace

import (
	"context"
	"sync"
)

type spanExporterCtxKey struct{}

type SpanExporter interface {
	ExportSpan(Span)
}

func ContextWithSpanExporter(ctx context.Context, spanExporter SpanExporter) context.Context {
	return context.WithValue(ctx, spanExporterCtxKey{}, spanExporter)
}

func SpanExporterFromContext(ctx context.Context) (SpanExporter, bool) {
	spanExporter, ok := ctx.Value(spanExporterCtxKey{}).(SpanExporter)

	return spanExporter, ok
}

type NoopSpanExporter struct{}

func (NoopSpanExporter) ExportSpan(Span) {}

type NonBlockingSpanExporter struct {
	lock     *sync.RWMutex
	exporter SpanExporter
}

func NewNonBlockingSpanExporter(wrappedExporter SpanExporter) NonBlockingSpanExporter {
	return NonBlockingSpanExporter{
		lock:     &sync.RWMutex{},
		exporter: wrappedExporter,
	}
}

func (e NonBlockingSpanExporter) ExportSpan(span Span) {
	e.lock.RLock()

	go func() {
		e.exporter.ExportSpan(span)

		e.lock.RUnlock()
	}()
}

func (e *NonBlockingSpanExporter) Close(ctx context.Context) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.exporter = NoopSpanExporter{}

	return nil
}
