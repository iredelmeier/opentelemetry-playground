module github.com/iredelmeier/opentelemetry-playground/bridges/opentracing

go 1.12

replace github.com/iredelmeier/opentelemetry-playground => ../..

require (
	github.com/iredelmeier/opentelemetry-playground v0.0.0-00010101000000-000000000000
	github.com/iredelmeier/opentelemetry-playground/headers v0.0.0-20190627035239-5a16137e1381
	github.com/lightstep/tracecontext.go v0.0.0-20181129014701-1757c391b1ac
	github.com/opentracing/opentracing-go v1.1.0
)
