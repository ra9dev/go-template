package http

import chi "github.com/go-chi/chi/v5"

// NewRouter constructor with middlewares for tracing
func NewRouter(params Params) chi.Router {
	router := chi.NewRouter()

	return DecorateHandler(router, params)
}
