package opentelemetry

import (
	"context"
	"sync"

	"github.com/iredelmeier/opentelemetry-playground/internal"
)

type Tracer struct {
	lock      *sync.RWMutex
	exporters []Exporter
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

func (t *Tracer) StartSpan(ctx context.Context, operationName string) context.Context {
	spanOpts := []internal.StartSpanOption{
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(t.finishSpan),
	}

	if parentID, ok := SpanIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithParentID(parentID))
	}

	if traceID, ok := TraceIDFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceID))
	}

	span := internal.NewSpan(spanOpts...)

	return internal.ContextWithSpan(ctx, span)
}

func (t *Tracer) finishSpan(span *internal.Span) {
	s := newSpan(span)

	for _, exporter := range t.exporters {
		t.lock.RLock()

		go func(e Exporter, span Span) {
			e.Export(span)

			t.lock.RUnlock()
		}(exporter, s)
	}
}
