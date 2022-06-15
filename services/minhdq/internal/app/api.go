package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"minhdq/internal/authentication"
	"minhdq/internal/service"
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

func NewRegisterServer() (*grpc.Server, error) {
	s := grpc.NewServer()
	regisS := service.GetRegisServer()
	authentication.RegisterResgisterServer(s, regisS)

	return s, nil
}

func NewLoginServer() (*grpc.Server, error) {
	s := grpc.NewServer()
	loginS := service.GetLoginServer()
	authentication.RegisterLoginServer(s, loginS)

	return s, nil
}
