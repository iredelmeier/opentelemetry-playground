package opentracing

import (
	"context"
	"time"

	"github.com/iredelmeier/opentelemetry-playground"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type Span struct {
	tracer *Tracer
	ctx    context.Context
}

func (s *Span) Finish() {
	s.FinishWithOptions(opentracing.FinishOptions{})
}

func (s *Span) FinishWithOptions(opts opentracing.FinishOptions) {
	opentelemetry.FinishSpan(s.ctx)
}

func (s *Span) Context() opentracing.SpanContext {
	spanContext := &SpanContext{}

	if id, ok := opentelemetry.SpanIDFromContext(s.ctx); ok {
		spanContext.id = id
	}

	if traceID, ok := opentelemetry.TraceIDFromContext(s.ctx); ok {
		spanContext.traceID = traceID
	}

	return spanContext
}

func (s *Span) SetOperationName(operationName string) opentracing.Span {
	// TODO
	return s
}

func (s *Span) SetTag(key string, value interface{}) opentracing.Span {
	// TODO
	return s
}

func (s *Span) LogFields(fields ...log.Field) {
	s.log(opentracing.LogRecord{
		Timestamp: time.Now(),
		Fields:    fields,
	})
}

func (s *Span) LogKV(alternatingKeyValues ...interface{}) {
	fields, err := log.InterleavedKVToFields(alternatingKeyValues...)
	if err != nil {
		s.LogFields(log.Error(err), log.String("function", "LogKV"))
		return
	}

	s.LogFields(fields...)
}

func (s *Span) SetBaggageItem(restrictedKey, value string) opentracing.Span {
	// TODO
	return s
}

func (s *Span) BaggageItem(restrictedKey string) string {
	// TODO
	return ""
}

func (s *Span) Tracer() opentracing.Tracer {
	return s.tracer
}

func (s *Span) LogEvent(event string) {
	s.LogEventWithPayload(event, nil)
}

func (s *Span) LogEventWithPayload(event string, payload interface{}) {
	s.Log(opentracing.LogData{
		Timestamp: time.Now(),
		Event:     event,
		Payload:   payload,
	})
}

func (s *Span) Log(logData opentracing.LogData) {
	s.log(logData.ToLogRecord())
}

func (s *Span) log(logRecord opentracing.LogRecord) {
	// TODO
}
