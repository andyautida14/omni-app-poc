package main

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	store := ds.CreateUserDatastore([]ds.User{
		{ID: "1", Name: "bob"},
	})
	usersH := handler.NewUsersHandler(store)
	userH := handler.NewUserHandler(store)
	mux.Handle("/users/", usersH)
	mux.Handle("/users/{id}", userH)
	http.ListenAndServe(":1337", mux)
}
