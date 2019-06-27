package internal

import (
	"sync"

	"github.com/gofrs/uuid"
)

type Span struct {
	id            [8]byte
	traceID       [16]byte
	parentID      [8]byte
	operationName string
	finishOnce    *sync.Once
	finishSpan    FinishSpan
}

func NewSpan(opts ...StartSpanOption) *Span {
	c := newStartSpanConfig(opts...)

	span := &Span{
		id:            c.id,
		traceID:       c.traceID,
		parentID:      c.parentID,
		operationName: c.operationName,
		finishOnce:    &sync.Once{},
		finishSpan:    c.finishSpan,
	}

	if span.id == [8]byte{} {
		u, _ := uuid.NewV4()

		copy(span.id[:], u[8:])
	}

	if span.traceID == [16]byte{} {
		u, _ := uuid.NewV4()

		span.traceID = u
	}

	return span
}

func (s *Span) OperationName() string {
	return s.operationName
}

func (s *Span) ID() [8]byte {
	return s.id
}

func (s *Span) TraceID() [16]byte {
	return s.traceID
}

func (s *Span) ParentID() [8]byte {
	return s.parentID
}

func (s *Span) Finish() {
	s.finishOnce.Do(func() {
		s.finishSpan(s)
	})
}
