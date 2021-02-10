package server

import (
	"net/http"
	uRoute "ohs-user/pkg/user/infrastructure/rest"

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
func New(userRt uRoute.UserRoute) Server {
	api := &Api{}
	r := mux.NewRouter()
	userRt.AddRoutes(r)
	api.router = r
	return api
}

func (a *Api) Router() http.Handler {
	return a.router
}
