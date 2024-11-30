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
	tmplLoader handler.HtmxTemplateLoader,
	staticFs http.FileSystem,
) {
	staticHandler := http.StripPrefix(STATIC_URL_PREFIX, http.FileServer(staticFs))
	mux.Handle(STATIC_URL_PREFIX, staticHandler)

	initRoute := handler.NewInitRouteFunc(tmplLoader, dsRegistry)

	mux.HandleFunc("/healthcheck/{$}", handler.HealthCheck)
	mux.Handle("/customers/new", initRoute(handler.Handlers{
		"GET": handler.RenderTemplate("main", []string{
			"shell",
			"customer-form",
		}),
	}))
	mux.Handle("/customers/{$}", initRoute(handler.Handlers{
		"POST": customer.SaveCustomer,
		"PUT":  customer.UpdateCustomer,
	}))
	mux.Handle("/customers/{id}/{$}", initRoute(handler.Handlers{
		"GET":    customer.Details("show"),
		"DELETE": customer.DeleteCustomer,
	}))
	mux.Handle("/customers/{id}/edit/{$}", initRoute(handler.Handlers{
		"GET": customer.Details("edit"),
	}))
	mux.Handle("/{$}", initRoute(handler.Handlers{
		"GET": home.GetHome,
	}))
}
