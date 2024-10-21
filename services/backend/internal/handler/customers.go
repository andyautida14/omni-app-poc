package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type customersHandler struct {
	tmplStore TemplateStore
	dataStore ds.CustomerDatastore
}

func (h *customersHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var c ds.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		handleInternalServerError(w, r)
		return
	}

	if err := h.dataStore.Create(c); err != nil {
		handleInternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *customersHandler) listCustomers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplStore.GetTemplates()
	if err != nil {
		handleInternalServerError(w, r)
		return
	}

	customers, err := h.dataStore.GetAll()
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

func NewCustomersHandler(tmplStore TemplateStore, dataStore ds.CustomerDatastore) http.Handler {
	return &customersHandler{tmplStore: tmplStore, dataStore: dataStore}
}
