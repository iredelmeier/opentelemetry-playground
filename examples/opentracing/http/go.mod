module github.com/iredelmeier/opentelemetry-playground/examples/opentracing/http

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../../..

replace github.com/iredelmeier/opentelemetry-playground/bridges/opentracing => ../../../bridges/opentracing

replace github.com/iredelmeier/opentelemetry-playground/exporters/file => ../../../exporters/file

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/bridges/opentracing v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/exporters/file v0.0.0-00010101000000-000000000000
	github.com/opentracing/opentracing-go v1.1.0
)
