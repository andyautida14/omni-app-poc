package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticDir embed.FS

func getStaticRootFs(staticPath string) (http.FileSystem, error) {
	if staticPath != "" {
		return http.Dir(staticPath), nil
	}

	staticFiles, err := fs.Sub(staticDir, "static")
	if err != nil {
		return nil, err
	}

	return http.FS(staticFiles), nil
}

func newStaticServer(urlPrefix string, fsPath string) (http.Handler, error) {
	rootFs, err := getStaticRootFs(fsPath)
	if err != nil {
		return nil, err
	}

	return http.StripPrefix(urlPrefix, http.FileServer(rootFs)), nil
}
