package server

import (
	aRoute "ms-asset/pkg/asset/infrastructure/rest"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	router http.Handler
}

// Server ...
type Server interface {
	Router() http.Handler
}

// New ...
func New(asset aRoute.AssetRoute) Server {
	api := &Api{}
	r := mux.NewRouter()
	asset.AddRoutes(r)
	api.router = r
	return api
}

func (a *Api) Router() http.Handler {
	return a.router
}
