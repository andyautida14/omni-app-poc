package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/andyautida/omni-app-poc/services/backend/internal/ds"
	"github.com/andyautida/omni-app-poc/services/backend/internal/handler"
)

//go:embed static/*
var staticDir embed.FS

func main() {
	users := []ds.User{
		{ID: "1", Name: "bob"},
	}
	store := ds.CreateUserDatastore(users)

	staticFiles, err := fs.Sub(staticDir, "static")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.FS(staticFiles))

	usersH := handler.NewUsersHandler(store)
	userH := handler.NewUserHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/users", usersH)
	mux.Handle("/users/", usersH)
	mux.Handle("/users/{id}", userH)
	mux.Handle("/users/{id}/", userH)
	mux.Handle("/", fs)
	http.ListenAndServe(":1337", mux)
}
