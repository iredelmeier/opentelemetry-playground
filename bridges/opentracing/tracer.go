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
	exporter opentelemetry.SpanExporter
}

func NewTracer(opts ...TracerOption) opentracing.Tracer {
	c := newTracerConfig(opts...)

	return &Tracer{
		exporter: c.exporter,
	}
}

func (t *Tracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	var config opentracing.StartSpanOptions

	for _, opt := range opts {
		opt.Apply(&config)
	}

	var sso []opentelemetry.StartSpanOption

	for _, ref := range config.References {
		if sc, ok := ref.ReferencedContext.(*SpanContext); ok {
			parentOpts := []opentelemetry.StartSpanOption{
				opentelemetry.WithTraceID(sc.traceID),
				opentelemetry.WithParentID(sc.id),
			}

			sso = append(sso, parentOpts...)

			break
		}
	}

	ctx := opentelemetry.ContextWithSpanExporter(context.Background(), t.exporter)

	ctx = opentelemetry.StartSpan(ctx, operationName, sso...)

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
