package router

import (
	"github.com/gorilla/mux"
	"github.com/rezaindrag/restapi/api"
	"github.com/rezaindrag/restapi/api/middlewares"
)

// New method initialize router
func New() *mux.Router {
	r := mux.NewRouter()

	// group routes
	newsRoutes := r.PathPrefix("/news").Subrouter()

	// middlewares
	middlewares.SetCommonMiddleware(r)

	// routers
	api.NewsRoutes(newsRoutes)

	return r
}
