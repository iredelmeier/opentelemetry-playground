module github.com/iredelmeier/opentelemetry-playground/examples/http

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../..

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/headers v0.0.0-00010101000000-000000000000
)

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/file => ../internal/exporters/file

replace github.com/iredelmeier/opentelemetry-playground/headers => ../../headers
