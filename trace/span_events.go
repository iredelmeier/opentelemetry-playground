package trace

import "time"

type StartSpanEvent struct {
	TraceContext  TraceContext
	ParentID      SpanID
	OperationName string
	StartTime     time.Time
}

type FinishSpanEvent struct {
	TraceContext TraceContext
	FinishTime   time.Time
	Attributes   map[string]string
}
