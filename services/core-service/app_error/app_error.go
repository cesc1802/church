package app_error

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int                    `json:"statusCode"`
	RootErr    error                  `json:"-"`
	Message    string                 `json:"message"`
	Log        string                 `json:"log"`
	ErrorKey   string                 `json:"errorKey"`
	VE         []ValidationErrorField `json:"ve,omitempty"`
}

type ValidationErrorField struct {
	Field        string `json:"field,omitempty"`
	Tag          string `json:"tag,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		ErrorKey:   key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		ErrorKey:   key,
	}
}

func NewUnAuthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		ErrorKey:   key,
	}
}

func ValidationError(msg string, key string, ve []ValidationErrorField) *AppError {
	appErr := NewErrorResponse(nil, msg, msg, key)
	appErr.VE = ve
	return appErr
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

var RecordNotFound = NewCustomError(nil, "record not found", ErrRecordNotFoundKey)

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), ErrInvalidRequestKey)
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"internal error", err.Error(), ErrInternalServerKey)
}

func ErrDBQuery(err error) *AppError {
	return NewCustomError(err, err.Error(), ErrDBQueryKey)
}
