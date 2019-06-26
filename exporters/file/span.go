package file

type Span struct {
	ID            string `json:"span_id"`
	TraceID       string `json:"trace_id"`
	OperationName string `json:"operation_name"`
}
