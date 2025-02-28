package router

import (
	"corason/internal/application/api/health"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	mainRouter := r.PathPrefix("/").Subrouter()
	apiV1Router := r.PathPrefix("/api/v1").Subrouter()

	mainRouter.HandleFunc("/health", health.NewHealth().Handle)
	apiV1Router.HandleFunc("/health", health.NewHealth().Handle)

	return r
}
