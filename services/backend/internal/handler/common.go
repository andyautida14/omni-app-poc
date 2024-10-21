package handler

import (
	"html/template"
	"net/http"
)

type TemplateStore interface {
	GetTemplates() (*template.Template, error)
}

func handleInternalServerError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
