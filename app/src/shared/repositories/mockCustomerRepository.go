package repositories

import (
	"errors"

	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type MockCustomerRepository struct {
	Customers []entity.Customer
}

func (r *MockCustomerRepository) Get(uuid uuid.UUID) (entity.Customer, error) {
	for i := 0; i < len(r.Customers); i++ {
		if r.Customers[i].Uuid.String() == uuid.String() {
			return r.Customers[i], nil
		}
	}
	return entity.Customer{}, errors.New("customer not found")
}

func (r *MockCustomerRepository) Update(customer entity.Customer) error {
	for i := 0; i < len(r.Customers); i++ {
		if r.Customers[i].Uuid.String() == customer.Uuid.String() {
			customer.SetUpdateDate()
			r.Customers[i] = customer
			return nil
		}
	}
	return errors.New("customer not found")
}

func (r *MockCustomerRepository) ListCustomers() ([]entity.Customer, error) {
	return r.Customers, nil
}

func (r *MockCustomerRepository) Create(customer entity.Customer) error {
	r.Customers = append(r.Customers, customer)
	return nil
}

func (r *MockCustomerRepository) Delete(uuid uuid.UUID) {
	var newCustomers []entity.Customer
	for i := 0; i < len(r.Customers); i++ {
		if r.Customers[i].Uuid == uuid {
			continue
		}

		newCustomers = append(newCustomers, r.Customers[i])
	}
	r.Customers = newCustomers
}
