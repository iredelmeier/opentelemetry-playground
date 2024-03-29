package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/examples/opentracing/http/internal"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	otTracer := bridge.NewTracer(bridge.WithExporter(exporter))

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	server := &http.Server{
		Addr:    internal.Host,
		Handler: traceHandler(otTracer, serveMux.ServeHTTP),
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func traceHandler(tracer opentracing.Tracer, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var startSpanOpts []opentracing.StartSpanOption

		if spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header)); err == nil {
			startSpanOpts = append(startSpanOpts, opentracing.ChildOf(spanContext))
		}

		operationName := fmt.Sprintf("HTTP GET: %s", r.URL.Path)
		span := tracer.StartSpan(operationName, startSpanOpts...)
		defer span.Finish()

		handler(w, r)
	}
}
