package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/iredelmeier/opentelemetry-playground/bridges/opencensus"
	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file"
	"github.com/iredelmeier/opentelemetry-playground/examples/opencensus/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/trace"
	"go.opencensus.io/plugin/ochttp"
	octrace "go.opencensus.io/trace"
)

func main() {
	exporter := trace.NewNonBlockingSpanExporter(file.NewExporter())
	defer exporter.Close(context.Background())

	ocExporter := opencensus.NewExporter(opencensus.WithSpanExporter(exporter))

	octrace.RegisterExporter(ocExporter)
	octrace.ApplyConfig(octrace.Config{DefaultSampler: octrace.AlwaysSample()})

	req := &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   internal.Host,
			Path:   internal.Path,
		},
	}

	req = req.WithContext(context.Background())

	client := &http.Client{
		Transport: &ochttp.Transport{},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	fmt.Printf("%s", body)
}
