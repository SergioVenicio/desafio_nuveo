package repositories

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type MockCustomerRepository struct {
	Customers []entity.Customer
}

func (r *MockCustomerRepository) ListCustomers() ([]entity.Customer, error) {
	return r.Customers, nil
}

func (r *MockCustomerRepository) Create(customer entity.Customer) error {
	r.Customers = append(r.Customers, customer)
	return nil
}
