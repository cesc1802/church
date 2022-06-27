package app

import (
	"context"
	"encoding/json"
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
	m.Config.PingPeriod = 5 * time.Second
	m.Config.PongWait = 5 * time.Second
	m.Config.WriteWait = 5 * time.Second
	return func(engine *gin.Engine) {

		engine.GET("/room", func(c *gin.Context) {
			data, err := service.RoomChatGetAll(c.Copy())
			if err != nil {
				c.Writer.Write([]byte(err.Error()))
				return
			}
			pl, _ := json.Marshal(data)
			c.Writer.Write(pl)
		})

		engine.GET("/room/:room", func(c *gin.Context) {
			room := c.Param("room")
			cmd := service.RoomChatIDCommand{ID: room}
			d, err := cmd.RoomChatGetOne(c.Copy())
			if err != nil {
				c.Writer.Write([]byte(err.Error()))
				return
			}
			pl, _ := json.Marshal(d)
			c.Writer.Write(pl)
		})

		engine.GET("/history/:room", func(c *gin.Context) {
			room := c.Param("room")
			cmd := service.RoomChatIDCommand{ID: room}
			d, err := cmd.RoomChatGetHistory(c.Copy())
			if err != nil {
				c.Writer.Write([]byte(err.Error()))
				return
			}
			pl, _ := json.Marshal(d)
			c.Writer.Write(pl)
		})

		engine.POST("/room/:room", func(c *gin.Context) {
			room := c.Param("room")
			cmd := service.RoomChatIDCommand{ID: room}
			err := cmd.NewRoom(c.Copy())
			if err != nil {
				c.Writer.Write([]byte(err.Error()))
				return
			}
			c.Writer.Write([]byte("Created room"))
		})

		engine.GET("/", func(c *gin.Context) {
			http.ServeFile(c.Writer, c.Request, "/home/minhdq/GolandProjects/church/services/minhdq/public/index.html")
		})

		engine.GET("/ws/:room/:name", func(c *gin.Context) {
			//secretKey := c.GetHeader("Authorization")
			//if secretKey != "123" {
			//	c.Writer.WriteHeader(http.StatusUnauthorized)
			//	c.Writer.Write([]byte("Wrong secret code"))
			//	return
			//}
			room := c.Param("room")
			name := c.Param("name")

			cmd := service.RoomChatIDCommand{ID: room}

			_, err := cmd.RoomChatGetOne(c.Copy())

			if err != nil {
				c.Writer.WriteHeader(400)
				c.Writer.Write([]byte(err.Error()))
				return
			}
			fmt.Println("User " + name + " connected to room " + room)
			m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"room": room, "name": name})
		})

		m.HandleConnect(func(session *melody.Session) {
			room, d := session.Get("room")
			if !d {
				session.Close()
				return
			}
			name, d := session.Get("name")
			if !d {
				session.Close()
				return
			}
			cmd := service.RoomChatIDPayloadCommand{
				ID:      room.(string),
				Payload: name.(string),
			}

			cmd.RoomChatAddUser(context.Background())
		})

		m.HandleError(func(session *melody.Session, err error) {
			if err != nil {
				room, d := session.Get("room")
				if !d {
					return
				}
				name, d := session.Get("name")
				if !d {
					session.Close()
					return
				}
				cmd := service.RoomChatIDPayloadCommand{
					ID:      room.(string),
					Payload: name.(string),
				}

				fmt.Println("closing the connection for " + name.(string) + " in room " + room.(string))

				cmd.RoomChatDeleteUser(context.Background())
				return
			}
		})

		m.HandleClose(func(session *melody.Session, i int, s string) error {
			room, d := session.Get("room")
			if !d {
				return nil
			}
			name, d := session.Get("name")
			if !d {
				session.Close()
				return nil
			}
			cmd := service.RoomChatIDPayloadCommand{
				ID:      room.(string),
				Payload: name.(string),
			}

			fmt.Println("closing the connection for " + name.(string) + " in room " + room.(string))

			err := cmd.RoomChatDeleteUser(context.Background())

			if err != nil {
				return err
			}

			return session.Close()
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
						name, d := session.Get("name")
						if !d {
							return false
						}
						cmd := service.RoomChatIDPayloadCommand{
							ID:      sr.(string),
							Payload: fmt.Sprintf("%s said: %s", name.(string), string(msg)),
						}

						cmd.RoomChatAddHistory(context.Background())
						return true
					}
				}
				return false
			})
		})
	}
}
