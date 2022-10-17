package http

import (
	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/ra9dev/go-template/docs" // swagger docs
)

type AdminAPI struct{}

func NewAdminAPI() AdminAPI {
	return AdminAPI{}
}

func (api AdminAPI) NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/v1/swagger/doc.json")))

	return router
}
