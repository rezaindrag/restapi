package api

import (
	"github.com/gorilla/mux"
	"github.com/rezaindrag/restapi/api/handlers"
)

// NewsRoutes contains news of routes
func NewsRoutes(s *mux.Router) {
	s.HandleFunc("", handlers.GetNews).Methods("GET")
	s.HandleFunc("/{id}", handlers.GetSingleNews).Methods("GET")
	s.HandleFunc("", handlers.StoreNews).Methods("POST")
	s.HandleFunc("/{id}", handlers.UpdateNews).Methods("PUT")
	s.HandleFunc("/{id}", handlers.DeleteNews).Methods("DELETE")
}
