package middleware

import (
	"github.com/gin-gonic/gin"
	"services.core-service/bool_utils"
	config "services.core-service/configs"
	"services.core-service/slice_utils"
)

func Cors(cfg *config.CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", slice_utils.SliceStringToString(cfg.AllowOrigins, ","))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", bool_utils.BoolToBoolString(cfg.AllowCredentials))
		c.Writer.Header().Set("Access-Control-Allow-Headers", slice_utils.SliceStringToString(cfg.AllowHeaders, ","))
		c.Writer.Header().Set("Access-Control-Allow-Methods", slice_utils.SliceStringToString(cfg.AllowMethods, ","))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
