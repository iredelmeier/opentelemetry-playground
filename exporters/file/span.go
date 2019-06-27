package file

import "time"

type Span struct {
	ID            string            `json:"span_id"`
	TraceID       string            `json:"trace_id"`
	ParentID      string            `json:"parent_id,omitempty"`
	OperationName string            `json:"operation_name"`
	StartTime     time.Time         `json:"start_time"`
	EndTime       time.Time         `json:"end_time"`
	Tags          map[string]string `json:"tags,omitempty"`
}
