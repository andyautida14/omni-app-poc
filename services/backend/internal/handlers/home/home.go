package home

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
)

type customerAllGetter interface {
	GetAll() ([]datastores.Customer, error)
}

func GetHome(
	tmplFactory handler.TemplateFactory,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	getTmpl := tmplFactory.CreateGetterFunc([]string{
		"shell",
		"customers",
	})
	customerDs := handler.DSMust[customerAllGetter](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := getTmpl()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		customers, err := customerDs.GetAll()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "shell", customers)
	}
}
