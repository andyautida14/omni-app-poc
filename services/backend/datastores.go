package main

import (
	"github.com/andyautida/omni-app-poc/lib/handler"
	"github.com/andyautida/omni-app-poc/services/backend/internal/datastores"
	"github.com/gocraft/dbr/v2"
)

// Register data stores here
func newDsRegistry(session *dbr.Session) handler.DatastoreRegistry {
	return handler.NewDsRegistry(
		datastores.NewCustomerDS(session),
	)
}
