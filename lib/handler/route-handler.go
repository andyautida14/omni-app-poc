package handler

import "net/http"

type routeHandler struct {
	handlers map[string]http.HandlerFunc
}

func (h *routeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := h.handlers[r.Method]
	if !ok {
		HandleNotFound(w, r)
		return
	}

	f(w, r)
}

func NewInitRouteFunc(
	tmplLoader HtmxTemplateLoader,
	dsRegistry DatastoreRegistry,
) func(Handlers) http.Handler {
	return func(handlerFactories Handlers) http.Handler {
		handlers := make(map[string]http.HandlerFunc)
		for method, initHandler := range handlerFactories {
			handlers[method] = initHandler(tmplLoader, dsRegistry)
		}
		return &routeHandler{handlers: handlers}
	}
}
