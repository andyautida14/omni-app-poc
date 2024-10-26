package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type customersHandler struct {
	dataStore ds.CustomerDatastore
}

func (h *customersHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var c ds.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		handleInternalServerError(w, r, err)
		return
	}

	if err := h.dataStore.Create(&c); err != nil {
		handleInternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *customersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		h.createCustomer(w, r)
	default:
		handleNotFound(w, r)
	}
}

func NewCustomersHandler(dataStore ds.CustomerDatastore) http.Handler {
	return &customersHandler{dataStore: dataStore}
}
