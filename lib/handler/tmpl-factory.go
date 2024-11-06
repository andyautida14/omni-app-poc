package handler

import (
	"html/template"
	"io/fs"
)

type tmplFactory struct {
	cacheFunc func(func() (*template.Template, error)) func() (*template.Template, error)
	fs.FS
}

func (f *tmplFactory) CreateGetterFunc(names []string) TemplateGetterFunc {
	getter := func() (*template.Template, error) {
		filenames := make([]string, len(names))
		for i, name := range names {
			filenames[i] = name + ".go.tmpl"
		}
		return template.ParseFS(f, filenames...)
	}

	return f.cacheFunc(getter)
}

func NewTmplFactory(fs fs.FS, templateCacheConfig string) (*tmplFactory, error) {
	cacheFunc, err := newCacheFunc[*template.Template](templateCacheConfig)
	if err != nil {
		return nil, err
	}

	return &tmplFactory{FS: fs, cacheFunc: cacheFunc}, nil
}
