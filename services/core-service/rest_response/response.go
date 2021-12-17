package rest_response

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"services.core-service/app_error"
	"services.core-service/i18n"
)

type JSONResponder interface {
	JSON(code int, obj interface{})
}

type Responder interface {
	SimpleSuccess(c JSONResponder, data interface{})
	Success(c JSONResponder, data, paging, filter interface{})
	Error(c JSONResponder, err error)
	SetLogTrace(logTrace bool)
}

func New(lang string, i18n *i18n.I18n) Responder {
	return &responder{
		lang: lang,
		i18n: i18n,
	}
}

type responder struct {
	lang string
	i18n *i18n.I18n
}

func (r responder) SimpleSuccess(c JSONResponder, data interface{}) {
	c.JSON(http.StatusOK, SimpleSuccessResponse(data))
}

func (r responder) Success(c JSONResponder, data, paging, filter interface{}) {
	c.JSON(http.StatusOK, NewSuccessResponse(data, paging, filter))
}

func (r responder) Error(c JSONResponder, err error) {
	if ve, ok := err.(validator.ValidationErrors); ok {
		appVE := app_error.HandleValidationErrors(r.lang, r.i18n, ve)
		c.JSON(appVE.StatusCode, appVE)
		return
	} else if appErr, ok := err.(*app_error.AppError); ok {
		c.JSON(appErr.StatusCode, app_error.HandleAppError(r.lang, r.i18n, appErr))
		return
	}
	c.JSON(http.StatusOK, err)
}

func (r responder) SetLogTrace(logTrace bool) {
	//panic("implement me")
}

type response struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SimpleSuccessResponse(data interface{}) *response {
	return NewSuccessResponse(data, nil, nil)
}

func NewSuccessResponse(data, paging, filter interface{}) *response {
	return &response{Data: data, Paging: paging, Filter: filter}
}

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	ErrorKey   string `json:"error_key"`
}

func NewErrorResponse(lang string, errorKey string, i18n *i18n.I18n) *errorResponse {
	message := i18n.MustLocalize(lang, errorKey, nil)
	return &errorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		ErrorKey:   errorKey,
	}
}
