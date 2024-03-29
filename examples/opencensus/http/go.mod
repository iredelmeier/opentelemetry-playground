module github.com/iredelmeier/opentelemetry-playground/examples/opencensus/http

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../../..

replace github.com/iredelmeier/opentelemetry-playground/bridges/opencensus => ../../../bridges/opencensus

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file => ../../internal/exporters/file

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/bridges/opencensus v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file v0.0.0-00010101000000-000000000000
	go.opencensus.io v0.22.0
)
