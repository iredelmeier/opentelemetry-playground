package file

import (
	"bytes"
	"encoding/hex"
	"encoding/json"

	"github.com/iredelmeier/opentelemetry-playground/trace"
)

type Exporter struct {
	encoder      *json.Encoder
	errorHandler ErrorHandler
}

func NewExporter(opts ...Option) Exporter {
	c := newConfig(opts...)

	return Exporter{
		encoder:      json.NewEncoder(c.file),
		errorHandler: c.errorHandler,
	}
}

func (e Exporter) ExportSpan(span trace.Span) {
	var parentID string
	if id := span.ParentID; !isEmptySpanID(id) {
		parentID = hex.EncodeToString(id[:])
	}

	if err := e.encoder.Encode(Span{
		ID:            hex.EncodeToString(span.TraceContext.SpanID[:]),
		TraceID:       hex.EncodeToString(span.TraceContext.TraceID[:]),
		ParentID:      parentID,
		OperationName: span.OperationName,
		StartTime:     span.StartTime,
		FinishTime:    span.FinishTime,
		Attributes:    span.Attributes,
	}); err != nil {
		e.errorHandler.Handle(err)
	}
}

func isEmptySpanID(spanID [8]byte) bool {
	var empty [8]byte

	return bytes.Equal(spanID[:], empty[:])
}
