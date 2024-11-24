package customer

import (
	"errors"
	"net/http"

	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
	"github.com/gocraft/dbr/v2"
)

type customerGetter interface {
	GetById(id string) (*datastores.Customer, error)
}

func GetDetails(
	tmplFactory handler.TemplateFactory,
	dsRegistry handler.DatastoreRegistry,
) http.HandlerFunc {
	getMainTmpl := tmplFactory.CreateGetterFunc([]string{
		"shell",
		"customer",
	})
	getNotFoundTmpl := tmplFactory.CreateGetterFunc([]string{
		"shell",
		"error-not-found",
	})
	customerDs := handler.DSMust[customerGetter](
		dsRegistry.Get("customer"),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		templateName := "shell"
		if r.Header.Get("HX-Request") == "true" {
			templateName = "main"
		}

		customer, err := customerDs.GetById(r.PathValue("id"))
		if err != nil {
			if errors.Is(err, dbr.ErrNotFound) {
				if tmpl, tmplErr := getNotFoundTmpl(); tmplErr != nil {
					err = tmplErr
				} else {
					w.WriteHeader(http.StatusNotFound)
					tmpl.ExecuteTemplate(w, templateName, nil)
					return
				}
			}

			handler.HandleInternalServerError(w, r, err)
			return
		}

		tmpl, err := getMainTmpl()
		if err != nil {
			handler.HandleInternalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, templateName, customer)
	}
}
