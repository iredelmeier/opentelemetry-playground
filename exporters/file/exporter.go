package file

import (
	"encoding/hex"
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
		ID:            hex.EncodeToString(spanData.ID[:]),
		TraceID:       hex.EncodeToString(spanData.TraceID[:]),
		OperationName: spanData.OperationName,
	}

	if err := e.encoder.Encode(span); err != nil {
		e.errorHandler.Handle(err)
	}
}
