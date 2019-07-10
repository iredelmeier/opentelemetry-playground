package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep"
	"github.com/iredelmeier/opentelemetry-playground/examples/lightstep/http/internal"
	"github.com/iredelmeier/opentelemetry-playground/headers"
	"github.com/iredelmeier/opentelemetry-playground/trace"
)

func main() {
	exporterOpts := []lightstep.Option{
		lightstep.WithComponentName("http-client"),
	}
	exporter, err := lightstep.NewExporter(exporterOpts...)
	if err != nil {
		panic(err)
	}
	defer exporter.Close(context.Background())

	req := &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   internal.Host,
			Path:   internal.Path,
		},
	}

	ctx := trace.ContextWithSpanExporter(req.Context(), exporter)

	ctx = trace.StartSpan(ctx, fmt.Sprintf("HTTP GET: %s", req.URL))

	trace.SetAttribute(ctx, "kind", "client")

	req = req.WithContext(ctx)

	injector := headers.NewInjector(req.Header)
	injector.Inject(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()

	trace.FinishSpan(ctx)

	fmt.Printf("%s", body)
}
