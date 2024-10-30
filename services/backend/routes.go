package main

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
	"github.com/gocraft/dbr/v2"
)

const STATIC_URL_PREFIX = "/static/"

func registerRoutes(mux *http.ServeMux, session *dbr.Session, c ServiceConfig) error {
	tmplFs, err := newTmplFs(c.TemplatePath, c.CacheTemplates)
	if err != nil {
		return err
	}

	staticFs, err := newStaticServer(STATIC_URL_PREFIX, c.StaticPath)
	if err != nil {
		return err
	}

	mux.HandleFunc("/healthcheck/", handler.HealthCheck)
	mux.Handle(STATIC_URL_PREFIX, staticFs)
	mux.Handle("/customers/", handler.NewCustomersHandler(tmplFs, session))
	mux.Handle("/", handler.NewHomeHandler(tmplFs, session))
	return nil
}
