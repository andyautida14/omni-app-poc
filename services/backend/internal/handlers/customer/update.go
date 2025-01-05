package customer

import (
	"encoding/json"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
)

type customerUpdater interface {
	UpdateOne(*datastores.Customer) error
}

func UpdateCustomer(
	_ handler.HtmxTemplateLoader,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	customerDs := handler.DSMust[customerUpdater](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		customer := &datastores.Customer{
			ID:        r.FormValue("id"),
			FirstName: r.FormValue("first-name"),
			LastName:  r.FormValue("last-name"),
		}

		if err := customerDs.UpdateOne(customer); err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		location, err := json.Marshal(map[string]string{
			"path":   "/customers/" + customer.ID + "/",
			"target": "#main",
		})
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
		}
		w.Header().Set("HX-Location", string(location))
		w.Header().Set("X-Notif-Msg", "Customer has been updated successfully.")
		w.Header().Set("X-Notif-Status", "danger")
		w.WriteHeader(http.StatusOK)
	}
}
