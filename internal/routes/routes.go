package routes

import (
	"github.com/go-chi/chi/v5"
)

type RouteRegister interface {
	RegisterRoutes(router *chi.Mux)
}
