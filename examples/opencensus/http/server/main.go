package main

import (
	"context"
	"io"
	"net/http"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/iredelmeier/opentelemetry-playground/bridges/opencensus"
	"github.com/iredelmeier/opentelemetry-playground/examples/opencensus/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	ocExporter := opencensus.NewExporter(opencensus.WithSpanExporter(exporter))

	trace.RegisterExporter(ocExporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	serveMux := http.NewServeMux()

	serveMux.HandleFunc(internal.Path, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	})

	handler := &ochttp.Handler{
		Handler: serveMux,
	}

	server := &http.Server{
		Addr:    internal.Host,
		Handler: handler,
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
