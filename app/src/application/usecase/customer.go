package usecase

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type CustomerUseCase struct {
	customers []entity.Customer
}

func (c *CustomerUseCase) List() []entity.Customer {
	return c.customers
}

func (c *CustomerUseCase) Create(name string, address string) (entity.Customer, error) {
	newCustomer, err := entity.NewCustomer(name, address)
	c.customers = append(c.customers, newCustomer)
	return newCustomer, err
}
