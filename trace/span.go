package trace

import (
	"context"
	"time"

	rootinternal "github.com/iredelmeier/opentelemetry-playground/internal"
	"github.com/iredelmeier/opentelemetry-playground/trace/internal"
)

type Span struct {
	TraceContext  TraceContext
	ParentID      SpanID
	OperationName string
	StartTime     time.Time
	FinishTime    time.Time
	Attributes    map[string]string
}

func StartSpan(ctx context.Context, operationName string, opts ...StartSpanOption) context.Context {
	c := newStartSpanConfig(opts...)

	spanOpts := []internal.StartSpanOption{
		internal.WithOperationName(operationName),
		internal.WithFinishSpan(finishSpan),
	}

	if traceContext, ok := TraceContextFromContext(ctx); ok {
		spanOpts = append(spanOpts, internal.WithTraceID(traceContext.TraceID))
		spanOpts = append(spanOpts, internal.WithParentID(traceContext.SpanID))
	}

	spanOpts = append(spanOpts, c.opts...)

	span := internal.NewSpan(spanOpts...)

	var spanEventExporter SpanEventExporter

	if exporter, ok := SpanEventExporterFromContext(ctx); ok {
		spanEventExporter = exporter
	}

	if spanEventExporter != nil {
		spanEventExporter.ExportStartSpanEvent(StartSpanEvent{
			TraceContext: TraceContext{
				TraceID: span.TraceID(),
				SpanID:  span.ID(),
			},
			ParentID:      span.ParentID(),
			OperationName: span.OperationName(),
			StartTime:     span.StartTime(),
		})
	}

	return internal.ContextWithSpan(ctx, span)
}

func SetAttribute(ctx context.Context, key string, value string) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		span.SetAttribute(key, value)
	}
}

func FinishSpan(ctx context.Context, opts ...FinishSpanOption) {
	if span, ok := internal.SpanFromContext(ctx); ok {
		c := newFinishSpanConfig(opts...)

		span.Finish(ctx, c.opts...)
	}
}

func finishSpan(ctx context.Context, span internal.Span) {
	s := Span{
		TraceContext: TraceContext{
			TraceID: span.TraceID(),
			SpanID:  span.ID(),
		},
		ParentID:      span.ParentID(),
		OperationName: span.OperationName(),
		StartTime:     span.StartTime(),
		FinishTime:    span.FinishTime(),
		Attributes:    make(map[string]string),
	}

	if attributes, ok := rootinternal.AttributesFromContext(ctx); ok {
		for _, attribute := range attributes.Entries() {
			s.Attributes[attribute.Key] = attribute.Value
		}
	}

	for _, attribute := range span.Attributes() {
		s.Attributes[attribute.Key] = attribute.Value
	}

	spanEventExporter := GlobalSpanEventExporter()

	if exporter, ok := SpanEventExporterFromContext(ctx); ok {
		spanEventExporter = exporter
	}

	if spanEventExporter != nil {
		spanEventExporter.ExportFinishSpanEvent(FinishSpanEvent{
			TraceContext: s.TraceContext,
			FinishTime:   s.FinishTime,
			Attributes:   s.Attributes,
		})
	}

	spanExporter := GlobalSpanExporter()

	if exporter, ok := SpanExporterFromContext(ctx); ok {
		spanExporter = exporter
	}

	if spanExporter != nil {
		spanExporter.ExportSpan(s)
	}
}
