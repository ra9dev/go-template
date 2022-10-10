package shutdown

import "errors"

var (
	ErrForceStop = errors.New("shutdown force stopped")
	ErrTimeout   = errors.New("shutdown timed out")
)
