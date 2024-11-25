package handler

import (
	"net/http"
)

type (
	HtmxTemplateLoader interface {
		Load([]string) (*HtmxTemplate, error)
	}

	DatastoreRegistry interface {
		Get(string) (interface{}, error)
	}

	Handlers map[string]HandlerFuncInit

	HandlerFuncInit func(HtmxTemplateLoader, DatastoreRegistry) http.HandlerFunc
)
