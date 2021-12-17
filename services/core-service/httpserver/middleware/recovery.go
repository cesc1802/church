package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"services.core-service/app_error"
	"services.core-service/httpserver/constants"
	"services.core-service/i18n"
)

func Recovery(i18n *i18n.I18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header(constants.HeaderContentType, "application/json")
				language := c.GetHeader(constants.HeaderAcceptLanguage)

				if ve, ok := err.(validator.ValidationErrors); ok {
					appVE := app_error.HandleValidationErrors(language, i18n, ve)
					c.AbortWithStatusJSON(appVE.StatusCode, appVE)
					return
				}

				if appErr, ok := err.(*app_error.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, app_error.HandleAppError(language, i18n, appErr))
					return
				}

				appErr := app_error.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
