package datastores

import (
	"time"

	"github.com/andyautida/omni-app-poc/lib/db"
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

type Customer struct {
	ID        string    `json:"-" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
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

func (ds *customerDatastore) RetrieveOne(id string) (*Customer, error) {
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

func (ds *customerDatastore) RetrieveMany(
	queryBuilder db.QueryBuilderFunc,
) ([]Customer, error) {
	var customers []Customer
	builder := ds.
		Select("*").
		From("customers")
	if _, err := queryBuilder(builder).
		Load(&customers); err != nil {
		return nil, err
	}
	return customers, nil
}

func NewCustomerDS(session *dbr.Session) func() (string, interface{}) {
	return func() (string, interface{}) {
		return "customer", &customerDatastore{Session: session}
	}
}
