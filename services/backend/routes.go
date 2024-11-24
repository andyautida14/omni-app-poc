package main

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handlers/customer"
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

	initRoute := handler.NewInitRouteFunc(tmplFactory, dsRegistry)

	mux.HandleFunc("/healthcheck/{$}", handler.HealthCheck)
	mux.Handle("/customers/new", initRoute(handler.Handlers{
		"GET": customer.NewCustomer,
	}))
	mux.Handle("/customers/{$}", initRoute(handler.Handlers{
		"POST": customer.SaveCustomer,
	}))
	mux.Handle("/customers/{id}/{$}", initRoute(handler.Handlers{
		"GET": customer.GetDetails,
	}))
	mux.Handle("/{$}", initRoute(handler.Handlers{
		"GET": home.GetHome,
	}))
}
