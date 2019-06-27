package headers

import (
	"context"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
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
	if parentSpan, ok := opentelemetry.ParentSpanFromContext(ctx); ok {
		traceParent := traceparent.TraceParent{
			Version: traceparent.Version,
			TraceID: parentSpan.TraceID,
			SpanID:  parentSpan.ID,
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
