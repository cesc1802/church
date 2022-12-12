package app

import (
	"encoding/json"
	"minhdq/internal/service"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	err = cmd.Register(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, nil)
}

func UserListAll(w http.ResponseWriter, r *http.Request) {
	data, err := service.FindAllUser(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}

func UserListID(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	data, err := cmd.FindByUserName(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}

func UserVertify(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	err = cmd.VerifyUser(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, nil)
}
