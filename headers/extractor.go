package headers

import (
	"context"
	"fmt"
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

func (e *Extractor) Extract(ctx context.Context) context.Context {
	traceContext, err := tracecontext.FromHeaders(e.headers)
	if err != nil {
		fmt.Println("ERR", err)
		return ctx
	}

	parentSpan := opentelemetry.ParentSpan{
		ID:      traceContext.TraceParent.SpanID,
		TraceID: traceContext.TraceParent.TraceID,
	}

	return opentelemetry.ContextWithParentSpan(ctx, parentSpan)
}
