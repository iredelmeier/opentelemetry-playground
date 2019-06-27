package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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

	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   internal.Host,
			Path:   internal.Path,
		},
	}
	ctx := tracer.StartSpan(context.Background(), fmt.Sprintf("HTTP GET: %s", req.URL))

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	opentelemetry.FinishSpan(ctx)

	fmt.Printf("%s", body)
}
