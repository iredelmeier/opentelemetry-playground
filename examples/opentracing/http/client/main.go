package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	bridge "github.com/iredelmeier/opentelemetry-playground/bridges/opentracing"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/examples/opentracing/http/internal"
	"github.com/opentracing/opentracing-go"
)

func main() {
	exporter := file.NewExporter()
	defer exporter.Close(context.Background())

	otTracer := bridge.NewTracer(bridge.WithExporter(exporter))

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
