package router

import (
	"github.com/gorilla/mux"
	"github.com/rezaindrag/restapi/api/handlers"
	"github.com/rezaindrag/restapi/api/middlewares"
)

// New method initialize router
func New() *mux.Router {
	r := mux.NewRouter()

	// middlewares
	r.Use(middlewares.CommonMiddleware)

	// routers
	r.HandleFunc("/news", handlers.GetNews).Methods("GET")
	r.HandleFunc("/news/{id}", handlers.GetSingleNews).Methods("GET")

	return r
}
