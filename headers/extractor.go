package headers

import (
	"context"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/lightstep/tracecontext.go"
)

type Extractor struct {
	headers http.Header
}

func NewExtractor(headers http.Header) *Extractor {
	return &Extractor{
		headers: headers,
	}
}

func (e *Extractor) Extract(ctx context.Context) opentelemetry.TraceContext {
	traceContext, err := tracecontext.FromHeaders(e.headers)
	if err != nil {
		return opentelemetry.TraceContext{}
	}

	return opentelemetry.TraceContext{
		TraceID: traceContext.TraceParent.TraceID,
		SpanID:  traceContext.TraceParent.SpanID,
	}
}
