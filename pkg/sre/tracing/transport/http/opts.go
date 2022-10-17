package http

import (
	trace "github.com/riandyrn/otelchi"
)

func buildOpts(handler Handler, params Params) []trace.Option {
	opts := []trace.Option{
		trace.WithChiRoutes(handler),
		// this is not necessary for vendors that properly implemented the tracing specs (e.g Jaeger, AWS X-Ray, etc...)
		trace.WithRequestMethodInSpanName(false),
	}

	opts = append(opts, params.ExtraOpts...)

	return opts
}
