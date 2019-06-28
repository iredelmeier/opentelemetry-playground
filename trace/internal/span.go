package internal

import (
	"context"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type Span struct {
	id            [8]byte
	traceID       [16]byte
	parentID      [8]byte
	operationName string
	startTime     time.Time
	finishTime    time.Time
	finishOnce    *sync.Once
	finishSpan    FinishSpan
}

func NewSpan(opts ...StartSpanOption) Span {
	c := newStartSpanConfig(opts...)

	span := Span{
		id:            c.id,
		traceID:       c.traceID,
		parentID:      c.parentID,
		operationName: c.operationName,
		startTime:     c.startTime,
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

func (s Span) OperationName() string {
	return s.operationName
}

func (s Span) ID() [8]byte {
	return s.id
}

func (s Span) TraceID() [16]byte {
	return s.traceID
}

func (s Span) ParentID() [8]byte {
	return s.parentID
}

func (s Span) StartTime() time.Time {
	return s.startTime
}

func (s Span) FinishTime() time.Time {
	return s.finishTime
}

func (s Span) Finish(ctx context.Context, opts ...FinishSpanOption) {
	s.finishOnce.Do(func() {
		c := newFinishSpanConfig(opts...)

		s.finishTime = c.finishTime

		s.finishSpan(ctx, s)
	})
}
