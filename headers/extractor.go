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

func (e *Extractor) Extract(ctx context.Context) opentelemetry.ParentSpan {
	traceContext, err := tracecontext.FromHeaders(e.headers)
	if err != nil {
		return opentelemetry.ParentSpan{}
	}

	return opentelemetry.ParentSpan{
		ID:      traceContext.TraceParent.SpanID,
		TraceID: traceContext.TraceParent.TraceID,
	}
}
