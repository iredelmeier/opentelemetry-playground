package opentracing

import (
	"context"

	"github.com/iredelmeier/opentelemetry-playground/trace"
	"github.com/opentracing/opentracing-go"
)

const (
	traceParentKey = "traceparent"
)

type Tracer struct {
	exporter trace.SpanExporter
}

func NewTracer(opts ...TracerOption) opentracing.Tracer {
	c := newTracerConfig(opts...)

	return Tracer{
		exporter: c.exporter,
	}
}

func (t Tracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	var config opentracing.StartSpanOptions

	for _, opt := range opts {
		opt.Apply(&config)
	}

	sso := []trace.StartSpanOption{
		trace.WithStartTime(config.StartTime),
	}

	for _, ref := range config.References {
		if sc, ok := ref.ReferencedContext.(*SpanContext); ok {
			parentOpts := []trace.StartSpanOption{
				trace.WithTraceID(sc.traceContext.TraceID),
				trace.WithParentID(sc.traceContext.SpanID),
			}

			sso = append(sso, parentOpts...)

			break
		}
	}

	ctx := trace.ContextWithSpanExporter(context.Background(), t.exporter)

	ctx = trace.StartSpan(ctx, operationName, sso...)

	return Span{
		tracer: t,
		ctx:    ctx,
	}
}

func (t Tracer) Inject(spanContext opentracing.SpanContext, format interface{}, carrier interface{}) error {
	switch sc := spanContext.(type) {
	case SpanContext:
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

func (t Tracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	switch format {
	case opentracing.Binary:
		return extractBinary(carrier)
	case opentracing.TextMap, opentracing.HTTPHeaders:
		return extractTextMap(carrier)
	default:
		return nil, opentracing.ErrUnsupportedFormat
	}
}
