package server

import (
	"backend/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {

	r := mux.NewRouter()

	routes.ConfigureRoutes(r)

	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      handler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
