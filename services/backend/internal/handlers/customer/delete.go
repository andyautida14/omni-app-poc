package customer

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
)

type customerDeleter interface {
	Delete(string) error
}

func DeleteCustomer(
	_ handler.HtmxTemplateLoader,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	customerDs := handler.DSMust[customerDeleter](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := customerDs.Delete(id); err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		w.Header().Set("HX-Location", `{"path":"/","target":"#main"}`)
		w.Header().Set("X-Notif-Msg", "Customer has been deleted successfully.")
		w.WriteHeader(http.StatusOK)
	}
}
