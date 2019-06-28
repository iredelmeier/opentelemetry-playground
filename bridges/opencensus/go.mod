module github.com/iredelmeier/opentelemetry-playground/bridges/opencensus

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../..

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	go.opencensus.io v0.22.0
)
