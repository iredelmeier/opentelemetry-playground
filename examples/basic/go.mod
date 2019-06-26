module github.com/iredelmeier/opentelemetry-playground/examples/basic

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../..

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/exporters/file v0.0.0-00010101000000-000000000000
)

replace github.com/iredelmeier/opentelemetry-playground/exporters/file => ../../exporters/file
