package usecase

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type CustomerUseCase struct {
	customers []entity.Customer
}

func (c *CustomerUseCase) List() []entity.Customer {
	var customers []entity.Customer

	customer, _ := entity.NewCustomer("Sergio", "Rua dalmacio de Azevedo, 233")
	c.customers = append(customers, customer)
	return c.customers
}

func (c *CustomerUseCase) Create(name string, address string) (entity.Customer, error) {
	newCustomer, err := entity.NewCustomer(name, address)
	c.customers = append(c.customers, newCustomer)
	return newCustomer, err
}
