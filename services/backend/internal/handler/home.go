package handler

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/gocraft/dbr/v2"
)

type homeHandler struct {
	tmplParser TemplateParser
	dataStore  ds.CustomerDatastore
}

func (h *homeHandler) getHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmplParser.ParseTemplates([]string{
		"shell",
		"customers",
	}, "get-home")
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

func NewHomeHandler(tmplParser TemplateParser, session *dbr.Session) http.Handler {
	h := &homeHandler{tmplParser: tmplParser, dataStore: ds.GetCustomerDS(session)}
	route := &routeHandler{handlers: map[string]http.HandlerFunc{
		"GET": h.getHome,
	}}
	return route
}
