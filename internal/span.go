package internal

import (
	"sync"
)

type Span struct {
	id            [8]byte
	traceID       [16]byte
	operationName string
	finishOnce    *sync.Once
	finishSpan    FinishSpan
}

func NewSpan(opts ...StartSpanOption) *Span {
	c := newStartSpanConfig(opts...)

	return &Span{
		id:            c.id,
		traceID:       c.traceID,
		operationName: c.operationName,
		finishOnce:    &sync.Once{},
		finishSpan:    c.finishSpan,
	}
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

func (s *Span) Finish() {
	s.finishOnce.Do(func() {
		s.finishSpan(s)
	})
}
