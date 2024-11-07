package handler

import (
	"html/template"
	"net/http"
)

type (
	TemplateFactory interface {
		CreateGetterFunc([]string) TemplateGetterFunc
	}

	TemplateGetterFunc func() (*template.Template, error)

	DatastoreRegistry interface {
		Get(string) (interface{}, error)
	}

	Handlers map[string]HandlerFuncInit

	HandlerFuncInit func(TemplateFactory, DatastoreRegistry) http.HandlerFunc
)
