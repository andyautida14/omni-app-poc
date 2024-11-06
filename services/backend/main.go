package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/db"
	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/sethvargo/go-envconfig"
)

var c ServiceConfig

//go:embed migrations/*
var migrationsDir embed.FS

//go:embed templates/*.go.tmpl
var embeddedTemplates embed.FS

//go:embed static/*
var staticDir embed.FS

func main() {
	ctx := context.Background()

	// load configuration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	// establish db connection
	conn, err := db.NewConn(c.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// perform db migration
	if err := conn.Migrate(migrationsDir); err != nil {
		log.Fatal(err)
	}

	// create datastores and datastore registry
	session := conn.NewSession(nil)
	dsRegistry := newDsRegistry(session)

	// get file system location for template files
	tmplFs, err := handler.GetTmplFilesFs(embeddedTemplates, c.TemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	// create template factory
	tmplFactory, err := handler.NewTmplFactory(tmplFs, c.TemplateCache)
	if err != nil {
		log.Fatal(err)
	}

	// get file system location for static files
	staticFs, err := handler.GetStaticFilesFs(staticDir, c.StaticPath)
	if err != nil {
		log.Fatal(err)
	}

	// create server and register routes
	mux := http.NewServeMux()
	registerRoutes(mux, dsRegistry, tmplFactory, staticFs)

	// run server
	if err := http.ListenAndServe(":1337", mux); err != nil {
		log.Fatal(err)
	}
}
