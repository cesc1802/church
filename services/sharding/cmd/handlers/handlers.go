package handlers

import (
	"net/http"
	"shard/module/auth_v1/transport/gin_auth"

	"github.com/gin-gonic/gin"

	core "services.core-service"
)

func EndUserRoutes(sc core.ServiceContext) func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		authV1 := engine.Group("/auth")
		{
			authV1.POST("/register", gin_auth.Register(sc))
			authV1.GET("/find/all", gin_auth.List(sc))
			authV1.GET("/find/:userid", gin_auth.FindByID(sc))
			authV1.GET("/fill", gin_auth.Fill(sc))
		}
	}
}
