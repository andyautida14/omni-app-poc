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
	isCacheEnabled bool
	tmplCache      map[string]*template.Template
	fs.FS
}

func (fs *tmplFs) parseTemplates(names []string) (*template.Template, error) {
	filenames := make([]string, len(names))
	for i, name := range names {
		filenames[i] = name + ".go.tmpl"
	}
	return template.ParseFS(fs, filenames...)
}

// TODO: remove
func (fs *tmplFs) ParseTemplates(names []string, cacheKey string) (*template.Template, error) {
	if !fs.isCacheEnabled {
		return fs.parseTemplates(names)
	}

	if fs.tmplCache == nil {
		fs.tmplCache = make(map[string]*template.Template)
	}

	tmpl, ok := fs.tmplCache[cacheKey]
	if ok {
		return tmpl, nil
	}

	tmpl, err := fs.parseTemplates(names)
	if err != nil {
		return nil, err
	}

	fs.tmplCache[cacheKey] = tmpl
	return tmpl, nil
}

// TODO: implement caching
func (fs *tmplFs) CreateGetterFunc(names []string) func() (*template.Template, error) {
	return func() (*template.Template, error) {
		return fs.parseTemplates(names)
	}
}

func newTmplFs(tmplPath string, isCacheEnabled bool) (*tmplFs, error) {
	if tmplPath != "" {
		fInfo, err := os.Stat(tmplPath)
		if err != nil {
			return nil, err
		}
		if !fInfo.IsDir() {
			return nil, fmt.Errorf("template path %s is not a directory", tmplPath)
		}
		return &tmplFs{FS: os.DirFS(tmplPath), isCacheEnabled: isCacheEnabled}, nil
	} else {
		tmplSubDir, err := fs.Sub(embeddedTemplates, "templates")
		if err != nil {
			return nil, err
		}
		return &tmplFs{FS: tmplSubDir, isCacheEnabled: isCacheEnabled}, nil
	}
}
