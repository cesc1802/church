package app

import (
	"github.com/go-chi/chi/v5"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func NewChiHandeler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.DefaultLogger)
		r.Use(middleware.Timeout(30 * time.Second))

		r.HandleFunc("/auth", AuthEndpoint)
		r.HandleFunc("/token", tokenEndpoint)

		r.Route("/user", func(r chi.Router) {
			r.Get("/", UserListAll)
			r.Get("/{userID}", UserListID)
			r.Post("/", UserCreate)
			r.Post("/verify", UserVertify)
		})

		r.Route("/client", func(r chi.Router) {
			r.Get("/", ClientListAll)
			r.Get("/{clientID}", ClientListID)
			r.Post("/", ClientCreate)
		})
	})

	return r
}
