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
	tmplLoader handler.HtmxTemplateLoader,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	tmpl := handler.TmplMust(tmplLoader.Load([]string{
		"shell",
		"customers",
	}))
	customerDs := handler.DSMust[customerAllGetter](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := customerDs.GetAll()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		tmpl.ExecuteHtmxTemplate(w, r, "main", customers)
	}
}
