package http

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	trace "github.com/riandyrn/otelchi"
)

// Handler is a type constraint for DecorateHandler
type Handler interface {
	chi.Routes
	Use(middlewares ...func(http.Handler) http.Handler)
}

// DecorateHandler with tracing for any MiddlewareUser
func DecorateHandler[HandlerType Handler](handler HandlerType, params Params) HandlerType {
	opts := buildOpts(handler, params)

	// TODO put request info middleware here as well
	handler.Use(trace.Middleware(params.Name, opts...))

	return handler
}
