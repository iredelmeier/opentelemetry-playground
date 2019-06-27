package opentelemetry

import (
	"context"
	"sync"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Tracer struct {
	lock      *sync.RWMutex
	exporters []SpanExporter
}

func NewTracer(opts ...TracerOption) *Tracer {
	c := newTracerConfig(opts...)

	tracer := &Tracer{
		lock:      &sync.RWMutex{},
		exporters: c.exporters,
	}

	return tracer
}

func (t *Tracer) Close(ctx context.Context) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.exporters = nil

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
	s := newSpan(ctx, span)

	for _, exporter := range t.exporters {
		t.lock.RLock()

		go func(e SpanExporter, span Span) {
			e.ExportSpan(span)

			t.lock.RUnlock()
		}(exporter, s)
	}
}
