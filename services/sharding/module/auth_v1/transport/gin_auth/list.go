package gin_auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	core "services.core-service"
	"services.core-service/app_error"
	"shard/constants"
	"shard/module/auth_v1/business"
	"shard/module/user_v1/storage"
)

func List(sc core.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := sc.MustGet(constants.KeyMainDb).(*gorm.DB)

		store := storage.NewPostgresUserStorage(db)
		biz := business.NewRegisterBusiness(store)
		data, err := biz.ListAll(c.Request.Context())
		if err != nil {
			app_error.MustError(err)
		}

		c.JSON(http.StatusOK, data)
	}
}
