package file

type Span struct {
	ID            string            `json:"span_id"`
	TraceID       string            `json:"trace_id"`
	ParentID      string            `json:"parent_id,omitempty"`
	OperationName string            `json:"operation_name"`
	Tags          map[string]string `json:"tags,omitempty"`
}
