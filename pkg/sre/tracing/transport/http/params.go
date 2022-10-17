package http

import (
	trace "github.com/riandyrn/otelchi"
)

// Params covers all possible tracing options
type Params struct {
	Name      string
	ExtraOpts []trace.Option
}

// NewParams constructor with minimal required fields
func NewParams(name string, extraOpts ...trace.Option) Params {
	return Params{
		Name:      name,
		ExtraOpts: extraOpts,
	}
}
