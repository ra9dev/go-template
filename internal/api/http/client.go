package http

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"

	_ "github.com/ra9dev/go-template/docs" // swagger docs
)

type ClientAPI struct{}

func NewClientAPI() ClientAPI {
	return ClientAPI{}
}

func (api ClientAPI) NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/ready", Handler(api.IsReady))
	router.Get("/live", Handler(api.IsLive))

	return router
}

func (api ClientAPI) IsReady(_ http.ResponseWriter, _ *http.Request) (struct{}, error) {
	return struct{}{}, nil
}

func (api ClientAPI) IsLive(_ http.ResponseWriter, _ *http.Request) (struct{}, error) {
	return struct{}{}, nil
}
