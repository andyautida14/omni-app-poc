package ds

import "github.com/gocraft/dbr/v2"

type Customer struct {
	ID        string
	FirstName string
	LastName  string
}

type CustomerDatastore interface {
	Create(Customer) error
	GetOne(string) (Customer, error)
	GetAll() ([]Customer, error)
}

type customerDatastore struct {
	session *dbr.Session
}

func (ds *customerDatastore) Create(c Customer) error {
	return nil
}

func (ds *customerDatastore) GetOne(id string) (Customer, error) {
	return Customer{}, nil
}

func (ds *customerDatastore) GetAll() ([]Customer, error) {
	return []Customer{}, nil
}

func NewCustomerDS(session *dbr.Session) CustomerDatastore {
	return &customerDatastore{session: session}
}
