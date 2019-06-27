package file

import (
	"bytes"
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

func (e *Exporter) ExportSpan(span opentelemetry.Span) {
	var parentID string
	if id := span.ParentID; !isEmptySpanID(id) {
		parentID = hex.EncodeToString(id[:])
	}

	if err := e.encoder.Encode(Span{
		ID:            hex.EncodeToString(span.ID[:]),
		TraceID:       hex.EncodeToString(span.TraceID[:]),
		ParentID:      parentID,
		OperationName: span.OperationName,
		Tags:          span.Tags,
	}); err != nil {
		e.errorHandler.Handle(err)
	}
}

func isEmptySpanID(spanID [8]byte) bool {
	var empty [8]byte

	return bytes.Equal(spanID[:], empty[:])
}
