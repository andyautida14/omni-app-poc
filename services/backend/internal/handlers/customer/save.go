package customer

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
)

type customerSaver interface {
	Save(*datastores.Customer) error
}

func SaveCustomer(
	_ handler.TemplateFactory,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	customerDs := handler.DSMust[customerSaver](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		customer := &datastores.Customer{
			FirstName: r.FormValue("first-name"),
			LastName:  r.FormValue("last-name"),
		}

		if err := customerDs.Save(customer); err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		w.Header().Set("HX-Location", `{"path":"/","target":"#main"}`)
		w.WriteHeader(201)
	}
}
