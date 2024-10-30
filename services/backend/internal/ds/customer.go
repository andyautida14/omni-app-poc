package ds

import (
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"-" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type CustomerDatastore interface {
	Create(*Customer) error
	GetOne(string) (Customer, error)
	GetAll() ([]Customer, error)
}

type customerDatastore struct {
	session *dbr.Session
}

var customerDsInstances map[*dbr.Session]CustomerDatastore

func (ds *customerDatastore) Create(c *Customer) error {
	now := time.Now()
	c.ID = uuid.New()
	c.CreatedAt = now
	c.UpdatedAt = now
	if _, err := ds.session.
		InsertInto("customers").
		Columns("id", "first_name", "last_name", "created_at", "updated_at").
		Record(c).
		Exec(); err != nil {
		return err
	}
	return nil
}

func (ds *customerDatastore) GetOne(id string) (Customer, error) {
	return Customer{}, nil
}

func (ds *customerDatastore) GetAll() ([]Customer, error) {
	var customers []Customer
	if _, err := ds.session.
		Select("*").
		From("customers").
		Load(&customers); err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomerDS(session *dbr.Session) CustomerDatastore {
	ds, ok := customerDsInstances[session]
	if !ok {
		ds = &customerDatastore{session: session}
	}
	return ds
}
