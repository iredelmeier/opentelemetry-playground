package headers

import (
	"context"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground/trace"
	"github.com/lightstep/tracecontext.go"
	"github.com/lightstep/tracecontext.go/traceparent"
)

type Injector struct {
	headers http.Header
}

func NewInjector(headers http.Header) *Injector {
	return &Injector{
		headers: headers,
	}
}

func (i *Injector) Inject(ctx context.Context) {
	if traceContext, ok := trace.TraceContextFromContext(ctx); ok {
		traceParent := traceparent.TraceParent{
			Version: traceparent.Version,
			TraceID: traceContext.TraceID,
			SpanID:  traceContext.SpanID,
			Flags: traceparent.Flags{
				Recorded: true,
			},
		}
		traceContext := tracecontext.TraceContext{
			TraceParent: traceParent,
		}

		traceContext.SetHeaders(i.headers)
	}
}
