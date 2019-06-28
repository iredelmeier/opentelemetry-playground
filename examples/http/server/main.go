package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground/examples/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/headers"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	exporter := trace.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	server := &http.Server{
		Addr:    internal.Host,
		Handler: traceHandler(exporter, serveMux.ServeHTTP),
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func traceHandler(exporter trace.SpanExporter, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractor := headers.NewExtractor(r.Header)
		parentSpan := extractor.Extract(r.Context())
		opts := []trace.StartSpanOption{
			trace.WithTraceID(parentSpan.TraceID),
			trace.WithParentID(parentSpan.SpanID),
		}

		ctx := trace.ContextWithSpanExporter(r.Context(), exporter)

		ctx = trace.StartSpan(ctx, fmt.Sprintf("HTTP GET: %s", r.URL.Path), opts...)
		defer trace.FinishSpan(ctx)

		r = r.WithContext(ctx)

		trace.SetAttribute(ctx, "kind", "server")

		handler(w, r)
	}
}
