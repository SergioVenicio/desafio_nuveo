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

func (r *MockCustomerRepository) Delete(uuid string) {
	var newCustomers []entity.Customer
	for i := 0; i < len(r.Customers)-1; i++ {
		if r.Customers[i].Uuid.String() == uuid {
			continue
		}

		newCustomers = append(newCustomers, r.Customers[i])
	}
	r.Customers = newCustomers
}
