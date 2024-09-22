package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type userHandler struct {
	store ds.UserDatastore
}

func (h *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		HandleNotFound(w, r)
		return
	}

	u, ok := h.store.GetOne(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		HandleInternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		h.getUser(w, r)
	}
}

func NewUserHandler(store ds.UserDatastore) http.Handler {
	return &userHandler{store: store}
}
