package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/reflection"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	core "services.core-service"
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
	reflection.Register(s)
	return s, nil
}

func NewLoginServer() (*grpc.Server, error) {
	s := grpc.NewServer()
	loginS := service.GetLoginServer()
	authentication.RegisterLoginServer(s, loginS)
	reflection.Register(s)
	return s, nil
}

func ChatRouter(sc core.ServiceContext, m *melody.Melody) func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.GET("/", func(c *gin.Context) {
			http.ServeFile(c.Writer, c.Request, "/home/minhdq/GolandProjects/church/services/minhdq/public/index.html")
		})

		engine.GET("/ws/:room", func(c *gin.Context) {
			room := c.Param("room")
			fmt.Println("Room " + room + " has now connections")
			m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"room": room})
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			room, d := s.Get("room")
			if !d {
				return
			}

			m.BroadcastFilter(msg, func(session *melody.Session) bool {
				sr, d := session.Get("room")
				if d {
					if sr == room {
						return true
					}
				}
				return false
			})
		})
	}
}
