package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type usersHandler struct {
	store ds.UserDatastore
}

func (h *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var u ds.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		HandleInternalServerError(w, r)
		return
	}

	h.store.Create(u)

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		HandleInternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (h *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	users := h.store.GetAll()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		HandleInternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *usersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		h.listUsers(w, r)
	case r.Method == http.MethodPost:
		h.createUser(w, r)
		return
	}
}

func NewUsersHandler(store ds.UserDatastore) http.Handler {
	return &usersHandler{store: store}
}
