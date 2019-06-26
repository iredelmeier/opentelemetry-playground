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

func (t *Tracer) StartSpan(operationName string) *Span {
	finishOnce := &sync.Once{}
	finish := func(span *Span) {
		finishOnce.Do(func() {
			t.finishSpan(span)
		})
	}

	return &Span{
		id:            internal.NewSpanID(),
		traceID:       internal.NewTraceID(),
		operationName: operationName,
		finish:        finish,
	}
}

func (t *Tracer) finishSpan(span *Span) {
	spanData := SpanData{
		ID:            span.id,
		TraceID:       span.traceID,
		OperationName: span.operationName,
	}

	for _, exporter := range t.exporters {
		t.lock.RLock()

		go func(e Exporter, sd SpanData) {
			e.Export(sd)

			t.lock.RUnlock()
		}(exporter, spanData)
	}
}
