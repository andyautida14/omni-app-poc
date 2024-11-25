package handler

import (
	"io/fs"
	"net/http"
	"text/template"
)

type HtmxTemplate struct {
	getTemplate func() (*template.Template, error)
}

type htmxTmplLoader struct {
	cacheFunc func(func() (*template.Template, error)) func() (*template.Template, error)
	fs.FS
}

func (t *HtmxTemplate) ExecuteHtmxTemplate(
	w http.ResponseWriter,
	r *http.Request,
	name string,
	data any,
) error {
	tmpl, err := t.getTemplate()
	if err != nil {
		return err
	}

	templateName := "shell"
	if r.Header.Get("HX-Request") == "true" {
		templateName = name
	}

	w.Header().Set("Content-Type", "text/html")
	return tmpl.ExecuteTemplate(w, templateName, data)
}

func (l *htmxTmplLoader) Load(names []string) (*HtmxTemplate, error) {
	getter := func() (*template.Template, error) {
		filenames := make([]string, len(names))
		for i, name := range names {
			filenames[i] = name + ".go.tmpl"
		}
		return template.ParseFS(l, filenames...)
	}

	return &HtmxTemplate{
		getTemplate: l.cacheFunc(getter),
	}, nil
}

func NewHtmxTmplLoader(fs fs.FS, templateCacheConfig string) (*htmxTmplLoader, error) {
	cacheFunc, err := newCacheFunc[*template.Template](templateCacheConfig)
	if err != nil {
		return nil, err
	}

	return &htmxTmplLoader{FS: fs, cacheFunc: cacheFunc}, nil
}
