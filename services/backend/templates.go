package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
)

//go:embed templates/*.go.tmpl
var embeddedTemplates embed.FS

type tmplFs struct {
	fs.FS
}

func (fs *tmplFs) ParseTemplates(names ...string) (*template.Template, error) {
	filenames := make([]string, len(names))
	for i, name := range names {
		filenames[i] = name + ".go.tmpl"
	}
	return template.ParseFS(fs, filenames...)
}

func newTmplFs(tmplPath string) (*tmplFs, error) {
	if tmplPath != "" {
		fInfo, err := os.Stat(tmplPath)
		if err != nil {
			return nil, err
		}
		if !fInfo.IsDir() {
			return nil, fmt.Errorf("template path %s is not a directory", tmplPath)
		}
		return &tmplFs{FS: os.DirFS(tmplPath)}, nil
	} else {
		tmplSubDir, err := fs.Sub(embeddedTemplates, "templates")
		if err != nil {
			return nil, err
		}
		return &tmplFs{FS: tmplSubDir}, nil
	}
}
