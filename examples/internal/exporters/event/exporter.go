package event

import (
	"sync"

	"github.com/iredelmeier/opentelemetry-playground/trace"
)

type Exporter struct {
	lock     *sync.Mutex
	spans    map[trace.TraceContext]trace.StartSpanEvent
	exporter trace.SpanExporter
}

func NewExporter(wrappedExporter trace.SpanExporter) Exporter {
	return Exporter{
		lock:     &sync.Mutex{},
		spans:    make(map[trace.TraceContext]trace.StartSpanEvent),
		exporter: wrappedExporter,
	}
}

func (e Exporter) ExportStartSpanEvent(event trace.StartSpanEvent) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.spans[event.TraceContext] = event
}

func (e Exporter) ExportFinishSpanEvent(event trace.FinishSpanEvent) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if start, ok := e.spans[event.TraceContext]; ok {
		delete(e.spans, event.TraceContext)

		span := trace.Span{
			TraceContext:  start.TraceContext,
			ParentID:      start.ParentID,
			OperationName: start.OperationName,
			StartTime:     start.StartTime,
			FinishTime:    event.FinishTime,
			Attributes:    event.Attributes,
		}

		e.exporter.ExportSpan(span)
	}
}
