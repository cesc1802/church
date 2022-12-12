package app

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"oauth/service"
)

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.ClientCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	secret, err := cmd.RegisterClient(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, map[string]string{"secret": secret})
}

func ClientListAll(w http.ResponseWriter, r *http.Request) {
	data, err := service.FindALlClients(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}

func ClientListID(w http.ResponseWriter, r *http.Request) {
	cmd := service.ClientCommand{}

	id := chi.URLParam(r, "clientID")

	if id == "" {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	cmd.ID = id

	data, err := cmd.FindClientById(r.Context())
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}
