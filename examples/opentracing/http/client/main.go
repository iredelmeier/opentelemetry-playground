package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/iredelmeier/opentelemetry-playground"
	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/examples/opentracing/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/exporters/file"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := opentelemetry.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	tracerOpts := []opentelemetry.TracerOption{
		opentelemetry.WithSpanExporter(exporter),
	}
	tracer := opentelemetry.NewTracer(tracerOpts...)

	otTracer := bridge.NewTracer(bridge.WithOpenTelemetryTracer(tracer))

	req := &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   internal.Host,
			Path:   internal.Path,
		},
	}

	operationName := fmt.Sprintf("HTTP GET: %s", req.URL)
	span := otTracer.StartSpan(operationName)

	if err := otTracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	span.Finish()

	fmt.Printf("%s", body)
}
