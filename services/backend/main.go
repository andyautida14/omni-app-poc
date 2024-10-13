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

	usersH := handler.NewUsersHandler(store)
	userH := handler.NewUserHandler(store)

	staticH, err := handler.NewStaticFilesHandler(false)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/users", usersH)
	mux.Handle("/users/", usersH)
	mux.Handle("/users/{id}", userH)
	mux.Handle("/users/{id}/", userH)
	mux.Handle("/", staticH)
	http.ListenAndServe(":1337", mux)
}
