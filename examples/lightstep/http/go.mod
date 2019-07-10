module github.com/iredelmeier/opentelemetry-playground/examples/lightstep/http

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../../..

replace github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep => ../../internal/exporters/lightstep

replace github.com/iredelmeier/opentelemetry-playground/headers => ../../../headers

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/examples/internal/exporters/lightstep v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/headers v0.0.0-00010101000000-000000000000
	github.com/lightstep/lightstep-tracer-common/golang/gogo v0.0.0-20190605223551-bc2310a04743 // indirect
	github.com/lightstep/lightstep-tracer-go v0.16.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/grpc v1.22.0 // indirect
)
