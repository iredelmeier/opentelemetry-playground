package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/examples/opentracing/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := file.NewExporter()
	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithSpanExporter(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)
	defer tracer.Close(context.Background())

	otTracer := bridge.NewTracer(bridge.WithOpenTelemetryTracer(tracer))

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
