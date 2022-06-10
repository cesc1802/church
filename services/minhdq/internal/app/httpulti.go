package app

import (
	"encoding/json"
	"net/http"
)

type ResponeErr struct {
	Message string `json:"message"`
	Code    int
}

var (
	ErrBadRequest = ResponeErr{
		Message: "Unable to process request",
		Code:    http.StatusBadRequest,
	}

	ErrUnProcessableEnity = ResponeErr{
		Message: "Unable to process entity",
		Code:    http.StatusUnprocessableEntity,
	}
)

func ResponeError(w http.ResponseWriter, Err ResponeErr) {
	w.WriteHeader(Err.Code)
	w.Write([]byte(Err.Message))
}

func ResponeData(w http.ResponseWriter, data interface{}) {
	respone, err := json.Marshal(data)

	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respone)
}
