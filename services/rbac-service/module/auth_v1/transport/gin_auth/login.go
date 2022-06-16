package gin_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	core "services.core-service"
	"services.core-service/app_error"
	"services.rbac-service/constants"
	"services.rbac-service/module/auth_v1/business"
	"services.rbac-service/module/auth_v1/dto"
	"services.rbac-service/module/user_v1/storage"
)

func Login(sc core.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginRequest

		if err := c.ShouldBind(&input); err != nil {
			app_error.MustError(app_error.ErrInvalidRequest(err))
		}
		db := sc.MustGet(constants.KeyMainDb).(*gorm.DB)

		store := storage.NewPostgresUserStorage(db)
		biz := business.NewLoginBusiness(store)
		res, err := biz.UserLogin(c.Request.Context(), &input)
		if err != nil {
			app_error.MustError(err)
		}

		c.JSON(http.StatusOK, &res)
	}
}
