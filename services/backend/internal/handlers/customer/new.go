package customer

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
)

func NewCustomer(
	tmplLoader handler.HtmxTemplateLoader,
	_ handler.DatastoreRegistry,
) http.HandlerFunc {
	tmpl := handler.TmplMust(tmplLoader.Load([]string{
		"shell",
		"customer-form",
	}))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteHtmxTemplate(w, r, "main", nil)
	}
}
