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
	if traceID, ok := trace.TraceIDFromContext(ctx); ok {
		if spanID, ok := trace.SpanIDFromContext(ctx); ok {
			traceParent := traceparent.TraceParent{
				Version: traceparent.Version,
				TraceID: traceID,
				SpanID:  spanID,
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
}
