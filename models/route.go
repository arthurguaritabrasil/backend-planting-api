package models

import "net/http"

type Route struct {
	Path         string
	Methods      []string
	Handler      http.HandlerFunc
	RequiredAuth bool
}
