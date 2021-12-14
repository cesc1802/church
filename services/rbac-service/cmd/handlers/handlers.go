package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	core "services.core-service"
	"services.rbac-service/module/auth_v1/transport/gin_auth"
)

func EndUserRoutes(sc core.ServiceContext) func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		authV1 := engine.Group("/auth")
		{
			authV1.POST("/register", gin_auth.Register(sc))
		}
	}
}
