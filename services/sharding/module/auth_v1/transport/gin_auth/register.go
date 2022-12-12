package gin_auth

import (
	"net/http"
	"shard/constants"

	"shard/module/auth_v1/business"
	"shard/module/auth_v1/dto"
	"shard/module/user_v1/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	core "services.core-service"
	"services.core-service/app_error"
)

func Register(sc core.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterRequest

		if err := c.ShouldBind(&input); err != nil {
			app_error.MustError(app_error.ErrInvalidRequest(err))
		}
		db := sc.MustGet(constants.KeyMainDb).(*gorm.DB)

		store := storage.NewPostgresUserStorage(db)
		biz := business.NewRegisterBusiness(store)

		if err := biz.Create(c.Request.Context(), &input); err != nil {
			app_error.MustError(err)
		}

		c.JSON(http.StatusOK, true)
	}
}
