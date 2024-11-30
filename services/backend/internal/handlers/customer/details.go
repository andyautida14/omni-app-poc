package customer

import (
	"errors"
	"log"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
	"github.com/gocraft/dbr/v2"
)

type customerGetter interface {
	GetById(id string) (*datastores.Customer, error)
}

func Details(page string) handler.HandlerFuncInit {
	var template string
	switch page {
	case "show":
		template = "customer"
	case "edit":
		template = "customer-form"
	default:
		log.Fatal("invalid details page type: ", page)
	}

	return func(
		tmplLoader handler.HtmxTemplateLoader,
		dsRegistry handler.DatastoreRegistry,
	) http.HandlerFunc {
		mainTmpl := handler.TmplMust(tmplLoader.Load([]string{
			"shell",
			template,
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
}
