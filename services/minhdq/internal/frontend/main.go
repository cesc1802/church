package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
)

func chat() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws/:room", func(c *gin.Context) {
		room := c.Param("room")
		fmt.Println("Room " + room + "has now connections")
		m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"room": room})
	})

	m.HandleConnect(func(session *melody.Session) {
		fmt.Println(session)
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

	r.Run(":5000")
}
