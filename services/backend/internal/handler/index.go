package handler

import (
	"html/template"
	"net/http"

	"embed"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
)

//go:embed templates/*.go.tmpl
var templates embed.FS

type indexHandler struct {
	store ds.UserDatastore
	tmpl  *template.Template
}

func (h *indexHandler) renderPage(w http.ResponseWriter, _ *http.Request) {
	users := h.store.GetAll()
	h.tmpl.ExecuteTemplate(w, "index.go.tmpl", users)
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		w.Header().Set("Content-type", "text/html")
		h.renderPage(w, r)
	}
}

func NewIndexHandler(store ds.UserDatastore) (http.Handler, error) {
	tmpl, err := template.ParseFS(templates, "templates/*")
	if err != nil {
		return nil, err
	}

	return &indexHandler{store: store, tmpl: tmpl}, nil
}
