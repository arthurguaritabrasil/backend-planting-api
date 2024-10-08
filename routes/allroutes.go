package routes

import "backend/models"

var Routes = map[string]models.Route{
	"register": {
		Path:         "/register",
		Methods:      []string{"POST"},
		Handler:      _,
		RequiredAuth: false,
	},
	"login": {
		Path:         "/login",
		Methods:      []string{"POST"},
		Handler:      _,
		RequiredAuth: false,
	},
}
