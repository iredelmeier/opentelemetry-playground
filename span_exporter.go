package opentelemetry

import (
	"context"
	"sync"
)

type SpanExporter interface {
	ExportSpan(Span)
}

type NoopSpanExporter struct{}

func (NoopSpanExporter) ExportSpan(Span) {}

type NonBlockingSpanExporter struct {
	lock     *sync.RWMutex
	exporter SpanExporter
}

func NewNonBlockingSpanExporter(wrappedExporter SpanExporter) *NonBlockingSpanExporter {
	return &NonBlockingSpanExporter{
		lock:     &sync.RWMutex{},
		exporter: wrappedExporter,
	}
}

func (e *NonBlockingSpanExporter) ExportSpan(span Span) {
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
