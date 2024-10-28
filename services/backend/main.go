package main

import (
	"context"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
	"github.com/sethvargo/go-envconfig"
)

const STATIC_URL_PREFIX = "/static/"

var c ServiceConfig

func main() {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	conn, err := openDbConn(c.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.migrate(); err != nil {
		log.Fatal(err)
	}

	customerStore := ds.NewCustomerDS(conn.NewSession(nil))
	customersH := handler.NewCustomersHandler(customerStore)

	tmplFs, err := newTmplFs(c.TemplatePath, c.CacheTemplates)
	if err != nil {
		log.Fatal(err)
	}
	homeH := handler.NewHomeHandler(tmplFs, customerStore)

	staticFs, err := newStaticServer(STATIC_URL_PREFIX, c.StaticPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle(STATIC_URL_PREFIX, staticFs)
	mux.Handle("/customers", customersH)
	mux.Handle("/customers/", customersH)
	mux.Handle("/", homeH)
	http.ListenAndServe(":1337", mux)
}
