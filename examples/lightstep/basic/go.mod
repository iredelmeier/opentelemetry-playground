module github.com/iredelmeier/opentelemetry-playground/examples/lightstep/basic

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../../..

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep => ../../internal/exporters/lightstep

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000 // indirect
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep v0.0.0-00010101000000-000000000000 // indirect
)
