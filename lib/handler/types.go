package handler

import "html/template"

type (
	TemplateFactory interface {
		CreateGetterFunc([]string) func() (*template.Template, error)
	}

	DatastoreRegistry interface {
		Get(string) (interface{}, error)
	}
)
