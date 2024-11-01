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

	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	conn, err := db.NewConn(c.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Migrate(migrationsDir); err != nil {
		log.Fatal(err)
	}

	session := conn.NewSession(nil)
	mux := http.NewServeMux()
	registerRoutes(mux, session, c)
	http.ListenAndServe(":1337", mux)
}
