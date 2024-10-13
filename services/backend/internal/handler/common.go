package handler

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed templates/*.go.tmpl
var embeddedTemplates embed.FS
var parsedTemplates *template.Template

func handleInternalServerError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func getTemplates() (*template.Template, error) {
	tmpl, err := template.ParseFS(embeddedTemplates, "templates/*")
	if err != nil {
		return nil, err
	}
	parsedTemplates = tmpl
	return parsedTemplates, nil
}
