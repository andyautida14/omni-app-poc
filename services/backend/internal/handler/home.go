package handler

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

type homeHandler struct {
	tmplParser TemplateParser
	dataStore  ds.CustomerDatastore
}

func (h *homeHandler) getHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplParser.ParseTemplates(
		"shell",
		"customers",
	)
	if err != nil {
		handleInternalServerError(w, r, err)
		return
	}

	customers, err := h.dataStore.GetAll()
	if err != nil {
		handleInternalServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "shell", customers)
}

func NewHomeHandler(tmplParser TemplateParser, dataStore ds.CustomerDatastore) http.Handler {
	h := &homeHandler{tmplParser: tmplParser, dataStore: dataStore}
	route := &routeHandler{handlers: map[string]http.HandlerFunc{
		"GET": h.getHome,
	}}
	return route
}
