package trace

import (
	"context"
	"sync"
)

type spanEventExporterCtxKey struct{}

type SpanEventExporter interface {
	ExportStartSpanEvent(StartSpanEvent)
	ExportFinishSpanEvent(FinishSpanEvent)
}

func ContextWithSpanEventExporter(ctx context.Context, spanEventExporter SpanEventExporter) context.Context {
	return context.WithValue(ctx, spanEventExporterCtxKey{}, spanEventExporter)
}

func SpanEventExporterFromContext(ctx context.Context) (SpanEventExporter, bool) {
	spanEventExporter, ok := ctx.Value(spanEventExporterCtxKey{}).(SpanEventExporter)

	return spanEventExporter, ok
}

type NoopSpanEventExporter struct{}

func (NoopSpanEventExporter) ExportStartSpanEvent(StartSpanEvent) {}

func (NoopSpanEventExporter) ExportFinishSpanEvent(FinishSpanEvent) {}

type NonBlockingSpanEventExporter struct {
	lock     *sync.RWMutex
	exporter SpanEventExporter
}

func NewNonBlockingSpanEventExporter(wrappedExporter SpanEventExporter) NonBlockingSpanEventExporter {
	return NonBlockingSpanEventExporter{
		lock:     &sync.RWMutex{},
		exporter: wrappedExporter,
	}
}

func (e NonBlockingSpanEventExporter) ExportStartSpanEvent(startSpanEvent StartSpanEvent) {
	e.lock.RLock()

	go func() {
		e.exporter.ExportStartSpanEvent(startSpanEvent)

		e.lock.RUnlock()
	}()
}

func (e NonBlockingSpanEventExporter) ExportFinishSpanEvent(finishSpanEvent FinishSpanEvent) {
	e.lock.RLock()

	go func() {
		e.exporter.ExportFinishSpanEvent(finishSpanEvent)

		e.lock.RUnlock()
	}()
}

func (e *NonBlockingSpanEventExporter) Close(ctx context.Context) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.exporter = NoopSpanEventExporter{}

	return nil
}
