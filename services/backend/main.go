package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/db"
	"github.com/sethvargo/go-envconfig"
)

//go:embed migrations/*
var migrationsDir embed.FS

var c ServiceConfig

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

	// create template registry
	tmplFs, err := newTmplFs(c.TemplatePath, c.CacheTemplates)
	if err != nil {
		log.Fatal(err)
	}

	// get file system location for static files
	staticFs, err := getStaticRootFs(c.StaticPath)
	if err != nil {
		log.Fatal(err)
	}

	// create server and register routes
	mux := http.NewServeMux()
	registerRoutes(mux, dsRegistry, tmplFs, staticFs)

	// run server
	if err := http.ListenAndServe(":1337", mux); err != nil {
		log.Fatal(err)
	}
}
