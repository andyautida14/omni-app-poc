package handler

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticDir embed.FS

func NewStaticFilesHandler() (http.Handler, error) {
	staticFiles, err := fs.Sub(staticDir, "static")
	if err != nil {
		return nil, err
	}

	return http.FileServer(http.FS(staticFiles)), nil
}
