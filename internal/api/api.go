package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewApi() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	
	r.Use(cors.Handler(cors.Options{
		// This is fine for now.
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},
		ExposedHeaders:   nil,
		AllowCredentials: true,
		MaxAge:           0,
	}))

	return r
}
