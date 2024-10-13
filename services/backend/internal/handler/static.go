package handler

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticDir embed.FS

func NewStaticFilesHandler(useEmbbedded bool) (http.Handler, error) {
	var staticRootFs http.FileSystem

	if useEmbbedded {
		staticFiles, err := fs.Sub(staticDir, "static")
		if err != nil {
			return nil, err
		}

		staticRootFs = http.FS(staticFiles)
	} else {
		staticRootFs = http.Dir("services/backend/internal/handler/static")
	}

	return http.FileServer(staticRootFs), nil
}
