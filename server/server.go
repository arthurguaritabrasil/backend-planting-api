package server

import (
	"log"
	"net/http"
	"time"
)

func StartServer() {

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
