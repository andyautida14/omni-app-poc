package main

import (
	"embed"
	"html/template"
	"path/filepath"
)

//go:embed templates/*.go.tmpl
var embeddedTemplates embed.FS

type tmplStore struct {
	tmplPathPattern string
	tmplCache       *template.Template
}

func (t *tmplStore) getEmbeddedTemplates() (*template.Template, error) {
	if t.tmplCache == nil {
		cache, err := template.ParseFS(embeddedTemplates, "templates/*")
		if err != nil {
			return nil, err
		}
		t.tmplCache = cache
	}

	return t.tmplCache, nil
}

func (t *tmplStore) getFsTemplates() (*template.Template, error) {
	return template.ParseGlob(t.tmplPathPattern)
}

func (t *tmplStore) GetTemplates() (*template.Template, error) {
	if t.tmplPathPattern != "" {
		return t.getFsTemplates()
	} else {
		return t.getEmbeddedTemplates()
	}
}

func newTmplStore(tmplPath string) *tmplStore {
	pattern := ""
	if tmplPath != "" {
		pattern = filepath.Join(tmplPath, "*.go.tmpl")
	}

	return &tmplStore{tmplPathPattern: pattern}
}
