package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type customersHandler struct {
	store ds.CustomerDatastore
}

func (h *customersHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var c ds.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		handleInternalServerError(w, r)
		return
	}

	if err := h.store.Create(c); err != nil {
		handleInternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *customersHandler) listCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.store.GetAll()
	if err != nil {
		handleInternalServerError(w, r)
		return
	}

	tmpl, err := getTemplates()
	if err != nil {
		handleInternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "customers.go.tmpl", customers)
}

func (h *customersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		h.listCustomers(w, r)
	case r.Method == http.MethodPost:
		h.createCustomer(w, r)
	default:
		handleNotFound(w, r)
	}
}

func NewCustomersHandler(store ds.CustomerDatastore) http.Handler {
	return &customersHandler{store: store}
}
