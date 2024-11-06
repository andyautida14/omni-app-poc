package handler

import "html/template"

type (
	TemplateFactory interface {
		CreateGetterFunc([]string) TemplateGetterFunc
	}

	TemplateGetterFunc func() (*template.Template, error)

	DatastoreRegistry interface {
		Get(string) (interface{}, error)
	}
)
