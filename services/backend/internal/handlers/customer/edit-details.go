package customer

import (
	"errors"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/gocraft/dbr/v2"
)

func EditDetails(
	tmplLoader handler.HtmxTemplateLoader,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	mainTmpl := handler.TmplMust(tmplLoader.Load([]string{
		"shell",
		"customer-form",
	}))
	notFoundTmpl := handler.TmplMust(tmplLoader.Load([]string{
		"shell",
		"error-not-found",
	}))
	customerDs := handler.DSMust[customerGetter](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		customer, err := customerDs.GetById(r.PathValue("id"))
		if err != nil {
			if errors.Is(err, dbr.ErrNotFound) {
				notFoundTmpl.ExecuteHtmxTemplate(w, r, "main", nil)
				return
			}

			handler.HandleInternalServerError(w, r, err)
			return
		}

		mainTmpl.ExecuteHtmxTemplate(w, r, "main", customer)
	}
}
