package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	for _, route := range Routes {
		var handler http.Handler = route.Handler
		if route.RequiredAuth {
			// Verificar a autenticacao dps
		}
		r.HandleFunc(route.Path, handler.ServeHTTP).Methods(route.Methods...)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path
	//
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	//
}
