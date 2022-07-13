package gin_auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	core "services.core-service"
	"shard/constants"
	"shard/module/user_v1/storage"
)

func Fill(sc core.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := sc.MustGet(constants.KeyMainDb).(*gorm.DB)

		store := storage.NewPostgresUserStorage(db)
		store.Fill(context.Background())

		c.JSON(http.StatusOK, nil)
	}
}
