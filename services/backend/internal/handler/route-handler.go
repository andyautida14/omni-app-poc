package handler

import "net/http"

type routeHandler struct {
	handlers map[string]http.HandlerFunc
}

func (h *routeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := h.handlers[r.Method]
	if !ok {
		handleNotFound(w, r)
		return
	}

	f(w, r)
}
