package sre

const (
	KeyRequestID Key = "request_id"
	KeyTraceID   Key = "trace_id"
	KeySpanID    Key = "span_id"
)

type Key string

func (k Key) String() string {
	return string(k)
}
