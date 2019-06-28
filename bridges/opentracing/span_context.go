package opentracing

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground/trace"
	"github.com/lightstep/tracecontext.go"
	"github.com/lightstep/tracecontext.go/traceparent"
	"github.com/opentracing/opentracing-go"
)

type SpanContext struct {
	id      trace.SpanID
	traceID trace.TraceID
}

func (sc *SpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

func (sc *SpanContext) traceParent() traceparent.TraceParent {
	return traceparent.TraceParent{
		Version: traceparent.Version,
		SpanID:  sc.id,
		TraceID: sc.traceID,
		Flags: traceparent.Flags{
			Recorded: true,
		},
	}
}

func (sc *SpanContext) injectBinary(carrier interface{}) error {
	switch c := carrier.(type) {
	case io.Writer:
		traceParent := sc.traceParent()

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
			TraceParent: sc.traceParent(),
		}

		traceContext.SetHeaders(c)

		return nil
	case opentracing.HTTPHeadersCarrier:
		return sc.injectTextMap(http.Header(c))
	case opentracing.TextMapWriter:
		c.Set(traceParentKey, sc.traceParent().String())

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

		return &SpanContext{
			id:      traceParent.SpanID,
			traceID: traceParent.TraceID,
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

		return &SpanContext{
			id:      traceContext.TraceParent.SpanID,
			traceID: traceContext.TraceParent.TraceID,
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

		return &SpanContext{
			id:      traceParent.SpanID,
			traceID: traceParent.TraceID,
		}, nil
	default:
		return nil, opentracing.ErrInvalidCarrier
	}
}
