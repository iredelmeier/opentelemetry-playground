package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/examples/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
)

func main() {
	exporter := file.NewExporter()
	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithExporters(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)
	defer tracer.Close(context.Background())

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	server := &http.Server{
		Addr:    internal.Host,
		Handler: traceHandler(tracer, serveMux.ServeHTTP),
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func traceHandler(tracer *opentelemetry.Tracer, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := tracer.StartSpan(r.Context(), fmt.Sprintf("HTTP GET: %s", r.URL.Path))
		defer opentelemetry.FinishSpan(ctx)

		r = r.WithContext(ctx)

		handler(w, r)
	}
}
