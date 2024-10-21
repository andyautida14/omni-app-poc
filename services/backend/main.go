package main

import (
	"context"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
	"github.com/sethvargo/go-envconfig"
)

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

	tmplStore := newTmplStore(c.TemplatePath)

	customerStore := ds.NewCustomerDS(conn.NewSession(nil))
	customersH := handler.NewCustomersHandler(tmplStore, customerStore)

	staticRootFs, err := getStaticRootFs(c.StaticPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/customers", customersH)
	mux.Handle("/customers/", customersH)
	mux.Handle("/", http.FileServer(staticRootFs))
	http.ListenAndServe(":1337", mux)
}
