package main

import (
	"embed"
	"net/url"

	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/lib/pq"
)

//go:embed migrations/*
var migrationsDir embed.FS

type db struct {
	u *url.URL
	*dbr.Connection
}

func (d *db) migrate() error {
	migrationSource, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(d.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", migrationSource, d.u.Path, driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func openDbConn(dbUrl string) (*db, error) {
	u, err := url.Parse(dbUrl)
	if err != nil {
		return nil, err
	}

	conn, err := dbr.Open(u.Scheme, dbUrl, nil)
	if err != nil {
		return nil, err
	}

	return &db{u: u, Connection: conn}, err
}
