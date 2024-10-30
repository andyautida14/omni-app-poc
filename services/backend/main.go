package main

import (
	"context"
	"log"
	"net/http"

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

	session := conn.NewSession(nil)
	mux := http.NewServeMux()
	registerRoutes(mux, session, c)
	http.ListenAndServe(":1337", mux)
}
