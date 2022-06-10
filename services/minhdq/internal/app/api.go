package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func NewChiHandeler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.DefaultLogger)
		r.Use(middleware.Timeout(30 * time.Second))

		r.Route("/user-group", func(r chi.Router) {
			r.Get("/", UserGroupAll)
			r.Post("/", UserGroupCreate)
			r.Delete("/{groupId}/{userId}", UserGroupDelete)
			r.Patch("/", UserGroupUpdate)
		})
	})

	return r
}
