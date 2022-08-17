package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/configs", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
			r.Put("/", func(w http.ResponseWriter, r *http.Request) {})
		})
	})
}
