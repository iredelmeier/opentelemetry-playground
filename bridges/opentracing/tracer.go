package opentracing

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/opentracing/opentracing-go"
)

const (
	traceParentKey = "traceparent"
)

type Tracer struct {
	tracer *opentelemetry.Tracer
}

func NewTracer(opts ...TracerOption) opentracing.Tracer {
	c := newTracerConfig(opts...)

	return &Tracer{
		tracer: c.openTelemetryTracer,
	}
}

func (t *Tracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	var config opentracing.StartSpanOptions

	for _, opt := range opts {
		opt.Apply(&config)
	}

	ctx := context.Background()

	for _, ref := range config.References {
		if sc, ok := ref.ReferencedContext.(*SpanContext); ok {
			var parentSpan opentelemetry.ParentSpan

			if id, ok := opentelemetry.SpanIDFromContext(sc.span.ctx); ok {
				parentSpan.ID = id
			}

			if traceID, ok := opentelemetry.TraceIDFromContext(sc.span.ctx); ok {
				parentSpan.TraceID = traceID
			}

			ctx = opentelemetry.ContextWithParentSpan(ctx, parentSpan)

			break
		}
	}

	ctx = t.tracer.StartSpan(ctx, operationName)

	return &Span{
		tracer: t,
		ctx:    ctx,
	}
}

func (t *Tracer) Inject(spanContext opentracing.SpanContext, format interface{}, carrier interface{}) error {
	switch sc := spanContext.(type) {
	case *SpanContext:
		switch format {
		case opentracing.Binary:
			return sc.injectBinary(carrier)
		case opentracing.TextMap, opentracing.HTTPHeaders:
			return sc.injectTextMap(carrier)
		default:
			return opentracing.ErrUnsupportedFormat
		}
	default:
		return opentracing.ErrSpanContextCorrupted
	}
}

func (t *Tracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	switch format {
	case opentracing.Binary:
		return extractBinary(carrier)
	case opentracing.TextMap, opentracing.HTTPHeaders:
		return extractTextMap(carrier)
	default:
		return nil, opentracing.ErrUnsupportedFormat
	}
}
