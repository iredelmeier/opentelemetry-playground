package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/examples/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/headers"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
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

func traceHandler(exporter opentelemetry.SpanExporter, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractor := headers.NewExtractor(r.Header)
		parentSpan := extractor.Extract(r.Context())
		opts := []opentelemetry.StartSpanOption{
			opentelemetry.WithTraceID(parentSpan.TraceID),
			opentelemetry.WithParentID(parentSpan.ID),
		}

		ctx := opentelemetry.ContextWithSpanExporter(r.Context(), exporter)

		ctx = opentelemetry.ContextWithKeyValue(ctx, "kind", "server")

		ctx = opentelemetry.StartSpan(ctx, fmt.Sprintf("HTTP GET: %s", r.URL.Path), opts...)
		defer opentelemetry.FinishSpan(ctx)

		r = r.WithContext(ctx)

		handler(w, r)
	}
}
