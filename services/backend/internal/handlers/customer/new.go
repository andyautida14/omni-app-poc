package customer

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
)

func NewCustomer(
	tmplFactory handler.TemplateFactory,
	_ handler.DatastoreRegistry,
) http.HandlerFunc {
	getTmpl := tmplFactory.CreateGetterFunc([]string{
		"shell",
		"customer-form",
	})

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := getTmpl()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		templateName := "shell"
		if r.Header.Get("HX-Request") == "true" {
			templateName = "main"
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, templateName, nil)
	}
}
