package admin

import (
	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/ra9dev/go-template/docs" // swagger docs
)

var swaggerHandler = httpSwagger.Handler(httpSwagger.URL("/v1/swagger/doc.json"))

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/swagger/*", swaggerHandler)

	return router
}
