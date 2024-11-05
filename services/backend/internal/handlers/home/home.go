package home

import (
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/db"
	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
)

type customerManyRetriever interface {
	RetrieveMany(db.QueryBuilderFunc) ([]datastores.Customer, error)
}

func GetHome(
	tmplFactory handler.TemplateFactory,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	getTmpl := tmplFactory.CreateGetterFunc([]string{
		"shell",
		"customers",
	})
	customerDs := handler.DSMust[customerManyRetriever](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := getTmpl()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		customers, err := customerDs.RetrieveMany(db.IdentityQueryBuilder)
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "shell", customers)
	}
}
