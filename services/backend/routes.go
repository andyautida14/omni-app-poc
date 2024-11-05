package main

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handlers/home"
)

const STATIC_URL_PREFIX = "/static/"

// Register HTTP routes and their handlers here
func registerRoutes(
	mux *http.ServeMux,
	dsRegistry handler.DatastoreRegistry,
	tmplFactory handler.TemplateFactory,
	staticFs http.FileSystem,
) {
	staticHandler := http.StripPrefix(STATIC_URL_PREFIX, http.FileServer(staticFs))
	mux.Handle(STATIC_URL_PREFIX, staticHandler)

	mux.HandleFunc("/healthcheck/{$}", handler.HealthCheck)
	mux.Handle("/{$}", handler.NewRouteHandler(map[string]http.HandlerFunc{
		"GET": home.GetHome(tmplFactory, dsRegistry),
	}))
}
