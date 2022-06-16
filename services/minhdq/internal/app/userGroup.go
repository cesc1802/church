package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"minhdq/internal/service"
)

func UserGroupUpdate(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserGroupUpdateCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	data, err := service.UpdateUserGroup(r.Context(), cmd)
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}

func UserGroupCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserGroupCreateCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	err = service.CreateUserGroup(r.Context(), cmd)

	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, nil)
}

func UserGroupAll(w http.ResponseWriter, r *http.Request) {
	data, err := service.UserGroupGetAll(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}

func UserGroupDelete(w http.ResponseWriter, r *http.Request) {
	cmd := service.UserGroupDeleteCommand{}

	groupID, err := strconv.Atoi(chi.URLParam(r, "groupId"))
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	userID, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	cmd.GroupID = groupID
	cmd.UserID = userID

	err = service.DeleteUserGroup(r.Context(), cmd)

	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, nil)
}
