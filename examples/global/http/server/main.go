package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground/examples/global/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/headers"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	trace.SetGlobalSpanExporter(exporter)

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	server := &http.Server{
		Addr:    internal.Host,
		Handler: traceHandler(serveMux.ServeHTTP),
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func traceHandler(handler http.HandlerFunc, opts ...Option) http.HandlerFunc {
	c := newConfig(opts...)

	return func(w http.ResponseWriter, r *http.Request) {
		parentSpan := c.extractor(r.Context(), r.Header)
		opts := []trace.StartSpanOption{
			trace.WithTraceID(parentSpan.TraceID),
			trace.WithParentID(parentSpan.SpanID),
		}

		ctx := trace.StartSpan(r.Context(), fmt.Sprintf("HTTP GET: %s", r.URL.Path), opts...)
		defer trace.FinishSpan(ctx)

		r = r.WithContext(ctx)

		trace.SetAttribute(ctx, "kind", "server")

		handler(w, r)
	}
}

type Extractor func(context.Context, http.Header) trace.TraceContext

func headerExtractor(ctx context.Context, header http.Header) trace.TraceContext {
	return headers.NewExtractor(header).Extract(ctx)
}

type Option func(*config)

type config struct {
	extractor Extractor
}

func WithExtractor(extractor Extractor) Option {
	return func(c *config) {
		c.extractor = extractor
	}
}

func newConfig(opts ...Option) config {
	var c config
	defaultOpts := []Option{
		WithExtractor(headerExtractor),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(&c)
	}

	return c
}
