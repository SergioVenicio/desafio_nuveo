package repositories

import "github.com/SergioVenicio/desafio_nuveo/domain/entity"

type ICustomerRepository interface {
	ListCustomers() ([]entity.Customer, error)
	Create(customer entity.Customer) error
	Delete(uuid string)
}
