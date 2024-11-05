package db

import "github.com/gocraft/dbr/v2"

type QueryBuilderFunc func(*dbr.SelectBuilder) *dbr.SelectBuilder
