package file

import (
	"encoding/json"

	"github.com/iredelmeier/opentelemetry-playground"
)

type Exporter struct {
	encoder      *json.Encoder
	errorHandler ErrorHandler
}

func NewExporter(opts ...Option) *Exporter {
	c := newConfig(opts...)

	return &Exporter{
		encoder:      json.NewEncoder(c.file),
		errorHandler: c.errorHandler,
	}
}

func (e *Exporter) Export(spanData opentelemetry.SpanData) {
	span := Span{
		OperationName: spanData.OperationName,
	}

	if err := e.encoder.Encode(span); err != nil {
		e.errorHandler.Handle(err)
	}
}
