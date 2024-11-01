package db

import (
	"io/fs"
	"net/url"

	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type dbConn struct {
	u *url.URL
	*dbr.Connection
}

func (db *dbConn) Migrate(migrationsDir fs.FS) error {
	migrationSource, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", migrationSource, db.u.Path, driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func NewConn(dbUrl string) (*dbConn, error) {
	u, err := url.Parse(dbUrl)
	if err != nil {
		return nil, err
	}

	conn, err := dbr.Open(u.Scheme, dbUrl, nil)
	if err != nil {
		return nil, err
	}

	return &dbConn{u: u, Connection: conn}, err
}
