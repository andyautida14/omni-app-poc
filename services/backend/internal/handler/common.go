package handler

import (
	"html/template"
	"log/slog"
	"net/http"
)

type TemplateParser interface {
	ParseTemplates([]string) (*template.Template, error)
}

func handleInternalServerError(w http.ResponseWriter, _ *http.Request, err error) {
	slog.Error("an unexpected error occured:" + err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
