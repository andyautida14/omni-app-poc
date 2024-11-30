package datastores

import (
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type Customer struct {
	ID        string    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type customerDatastore struct {
	*dbr.Session
}

func (ds *customerDatastore) Save(c *Customer) error {
	now := time.Now()
	c.ID = uuid.New().String()
	c.CreatedAt = now
	c.UpdatedAt = now
	if _, err := ds.
		InsertInto("customers").
		Columns("id", "first_name", "last_name", "created_at", "updated_at").
		Record(c).
		Exec(); err != nil {
		return err
	}
	return nil
}

func (ds *customerDatastore) GetById(id string) (*Customer, error) {
	customer := &Customer{}
	if err := ds.
		Select("*").
		From("customers").
		Where("id = ?", id).
		LoadOne(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (ds *customerDatastore) GetAll() ([]Customer, error) {
	var customers []Customer
	if _, err := ds.
		Select("*").
		From("customers").
		Load(&customers); err != nil {
		return nil, err
	}
	return customers, nil
}

func (ds *customerDatastore) UpdateOne(c *Customer) error {
	now := time.Now()
	res, err := ds.
		Update("customers").
		Set("first_name", c.FirstName).
		Set("last_name", c.LastName).
		Set("updated_at", now).
		Where("id = ?", c.ID).
		Exec()
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return dbr.ErrNotFound
	}
	c.UpdatedAt = now
	return nil
}

func (ds *customerDatastore) Delete(id string) error {
	_, err := ds.DeleteFrom("customers").
		Where("id = ?", id).
		Exec()
	return err
}

func NewCustomerDS(session *dbr.Session) func() (string, interface{}) {
	return func() (string, interface{}) {
		return "customer", &customerDatastore{Session: session}
	}
}
