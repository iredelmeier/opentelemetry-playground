package opentracing

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/lightstep/tracecontext.go"
	"github.com/lightstep/tracecontext.go/traceparent"
	"github.com/opentracing/opentracing-go"
)

type SpanContext struct {
	span *Span
}

func (sc *SpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

func (sc *SpanContext) injectBinary(carrier interface{}) error {
	switch c := carrier.(type) {
	case io.Writer:
		traceParent := sc.span.traceParent()

		if _, err := c.Write([]byte(traceParent.String())); err != nil {
			return err
		}

		return nil
	default:
		return opentracing.ErrInvalidCarrier
	}
}

func (sc *SpanContext) injectTextMap(carrier interface{}) error {
	switch c := carrier.(type) {
	case http.Header:
		traceContext := tracecontext.TraceContext{
			TraceParent: sc.span.traceParent(),
		}

		traceContext.SetHeaders(c)

		return nil
	case opentracing.HTTPHeadersCarrier:
		return sc.injectTextMap(http.Header(c))
	case opentracing.TextMapWriter:
		c.Set(traceParentKey, sc.span.traceParent().String())

		return nil
	default:
		return opentracing.ErrInvalidCarrier
	}
}

func extractBinary(carrier interface{}) (*SpanContext, error) {
	switch c := carrier.(type) {
	case io.Reader:
		b, err := ioutil.ReadAll(c)
		if err != nil {
			return nil, err
		}

		traceParent, err := traceparent.Parse(b)
		if err != nil {
			return nil, err
		}

		parentSpan := opentelemetry.ParentSpan{
			ID:      traceParent.SpanID,
			TraceID: traceParent.TraceID,
		}

		return &SpanContext{
			span: &Span{
				ctx: opentelemetry.ContextWithParentSpan(context.Background(), parentSpan),
			},
		}, nil
	default:
		return nil, opentracing.ErrInvalidCarrier
	}
}

func extractTextMap(carrier interface{}) (*SpanContext, error) {
	switch c := carrier.(type) {
	case http.Header:
		traceContext, err := tracecontext.FromHeaders(c)
		if err != nil {
			return nil, err
		}

		parentSpan := opentelemetry.ParentSpan{
			ID:      traceContext.TraceParent.SpanID,
			TraceID: traceContext.TraceParent.TraceID,
		}

		return &SpanContext{
			span: &Span{
				ctx: opentelemetry.ContextWithParentSpan(context.Background(), parentSpan),
			},
		}, nil
	case opentracing.HTTPHeadersCarrier:
		return extractTextMap(http.Header(c))
	case opentracing.TextMapReader:
		var traceParent traceparent.TraceParent

		if err := c.ForeachKey(func(key string, value string) error {
			if key == traceParentKey {
				var err error
				traceParent, err = traceparent.ParseString(value)
				if err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		parentSpan := opentelemetry.ParentSpan{
			ID:      traceParent.SpanID,
			TraceID: traceParent.TraceID,
		}

		return &SpanContext{
			span: &Span{
				ctx: opentelemetry.ContextWithParentSpan(context.Background(), parentSpan),
			},
		}, nil
	default:
		return nil, opentracing.ErrInvalidCarrier
	}
}
