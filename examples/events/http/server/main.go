package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground/examples/events/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/event"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/headers"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	spanExporter := file.NewExporter()
	defer spanExporter.Close(context.Background())

	eventExporter := event.NewExporter(spanExporter)

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	server := &http.Server{
		Addr:    internal.Host,
		Handler: traceHandler(eventExporter, serveMux.ServeHTTP),
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func traceHandler(exporter trace.SpanEventExporter, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractor := headers.NewExtractor(r.Header)
		parentSpan := extractor.Extract(r.Context())
		opts := []trace.StartSpanOption{
			trace.WithTraceID(parentSpan.TraceID),
			trace.WithParentID(parentSpan.SpanID),
		}

		ctx := trace.ContextWithSpanEventExporter(r.Context(), exporter)

		ctx = trace.StartSpan(ctx, fmt.Sprintf("HTTP GET: %s", r.URL.Path), opts...)
		defer trace.FinishSpan(ctx)

		r = r.WithContext(ctx)

		trace.SetAttribute(ctx, "kind", "server")

		handler(w, r)
	}
}
