package main

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
)

func main() {
	users := []ds.User{
		{ID: "1", Name: "bob"},
	}
	store := ds.CreateUserDatastore(users)

	indexH, err := handler.NewIndexHandler(store)
	if err != nil {
		panic(err)
	}
	usersH := handler.NewUsersHandler(store)
	userH := handler.NewUserHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/", indexH)
	mux.Handle("/users", usersH)
	mux.Handle("/users/", usersH)
	mux.Handle("/users/{id}", userH)
	mux.Handle("/users/{id}/", userH)
	http.ListenAndServe(":1337", mux)
}
