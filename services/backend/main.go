package main

import (
	"context"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
	"github.com/gocraft/dbr/v2"
	"github.com/sethvargo/go-envconfig"

	_ "github.com/lib/pq"
)

type ServiceConfig struct {
	DbDriver string `env:"DB_DRIVER, default=postgres"`
	DbUrl    string `env:"DB_URL, required"`
}

func main() {
	ctx := context.Background()

	var c ServiceConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	conn, err := dbr.Open(c.DbDriver, c.DbUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	customerStore := ds.NewCustomerDS(*conn)
	customersH := handler.NewCustomersHandler(customerStore)

	users := []ds.User{
		{ID: "1", Name: "bob"},
	}
	store := ds.CreateUserDatastore(users)

	usersH := handler.NewUsersHandler(store)
	userH := handler.NewUserHandler(store)

	staticH, err := handler.NewStaticFilesHandler(false)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/customers", customersH)
	mux.Handle("/customers/", customersH)
	mux.Handle("/users", usersH)
	mux.Handle("/users/", usersH)
	mux.Handle("/users/{id}", userH)
	mux.Handle("/users/{id}/", userH)
	mux.Handle("/", staticH)
	http.ListenAndServe(":1337", mux)
}
