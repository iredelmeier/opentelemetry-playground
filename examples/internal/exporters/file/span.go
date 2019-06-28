package file

import "time"

type Span struct {
	ID            string            `json:"span_id"`
	TraceID       string            `json:"trace_id"`
	ParentID      string            `json:"parent_id,omitempty"`
	OperationName string            `json:"operation_name"`
	StartTime     time.Time         `json:"start_time"`
	FinishTime    time.Time         `json:"finish_time"`
	Attributes    map[string]string `json:"attributes,omitempty"`
}
