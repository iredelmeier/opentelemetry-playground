package lightstep

import (
	"context"
	"encoding/binary"
	"errors"

	"github.com/iredelmeier/opentelemetry-playground/trace"
	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing/opentracing-go"
)

var (
	ErrFailedToCreateExporter = errors.New("lightstep: failed to create exporter")
)

type Exporter struct {
	tracer lightstep.Tracer
}

func NewExporter(opts ...Option) (Exporter, error) {
	c := newConfig(opts...)
	tracer := lightstep.NewTracer(c.tracerOptions)

	if tracer == nil {
		if err := c.tracerOptions.Validate(); err != nil {
			return Exporter{}, err
		}
		return Exporter{}, ErrFailedToCreateExporter
	}

	return Exporter{
		tracer: tracer,
	}, nil
}

func (e Exporter) ExportSpan(span trace.Span) {
	traceID := binary.BigEndian.Uint64(span.TraceContext.TraceID[8:])
	spanID := binary.BigEndian.Uint64(span.TraceContext.SpanID[:])
	parentID := binary.BigEndian.Uint64(span.ParentID[:])
	startSpanOpts := []opentracing.StartSpanOption{
		opentracing.StartTime(span.StartTime),
		lightstep.SetTraceID(traceID),
		lightstep.SetSpanID(spanID),
		lightstep.SetParentSpanID(parentID),
	}

	otSpan := e.tracer.StartSpan(span.OperationName, startSpanOpts...)

	for k, v := range span.Attributes {
		otSpan.SetTag(k, v)
	}

	finishOpts := opentracing.FinishOptions{
		FinishTime: span.FinishTime,
	}

	otSpan.FinishWithOptions(finishOpts)
}

func (e Exporter) Close(ctx context.Context) error {
	e.tracer.Close(ctx)

	return nil
}
