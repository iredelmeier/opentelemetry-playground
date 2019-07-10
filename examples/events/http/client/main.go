package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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

	req := &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "http",
			Host:   internal.Host,
			Path:   internal.Path,
		},
	}

	ctx := trace.ContextWithSpanEventExporter(req.Context(), eventExporter)

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
