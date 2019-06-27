package opentelemetry

import (
	"context"
	"sync"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Tracer struct {
	lock     *sync.RWMutex
	exporter SpanExporter
}

func NewTracer(opts ...TracerOption) *Tracer {
	c := newTracerConfig(opts...)

	return &Tracer{
		lock:     &sync.RWMutex{},
		exporter: c.exporter,
	}
}

func (t *Tracer) Close(ctx context.Context) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.exporter = NoopSpanExporter{}

	return nil
}

func (t *Tracer) StartSpan(ctx context.Context, operationName string, opts ...StartSpanOption) context.Context {
	c := newStartSpanConfig(opts...)

	spanOpts := []internal.StartSpanOption{
		internal.WithID(c.id),
		internal.WithTraceID(c.traceID),
		internal.WithParentID(c.parentID),
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(t.finishSpan),
	}

	if traceID, ok := TraceIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceID))
	}

	if parentID, ok := SpanIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithParentID(parentID))
	}

	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}

func (t *Tracer) finishSpan(ctx context.Context, span *internal.Span) {
	t.lock.RLock()

	go func() {
		defer t.lock.RUnlock()

		t.exporter.ExportSpan(newSpan(ctx, span))
	}()
}
