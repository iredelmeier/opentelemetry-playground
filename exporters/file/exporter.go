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

func (e *Exporter) Export(span opentelemetry.Span) {
	if err := e.encoder.Encode(Span{
		ID:            hex.EncodeToString(span.ID[:]),
		TraceID:       hex.EncodeToString(span.TraceID[:]),
		OperationName: span.OperationName,
	}); err != nil {
		e.errorHandler.Handle(err)
	}
}
