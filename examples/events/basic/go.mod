module github.com/iredelmeier/opentelemetry-playground/examples/events/basic

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../../..

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/event => ../../internal/exporters/event

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file => ../../internal/exporters/file

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/bridges/opentracing v0.0.0-20190710190858-da7ae7385f7a
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/event v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file v0.0.0-00010101000000-000000000000
	github.com/opentracing/opentracing-go v1.1.0
)
