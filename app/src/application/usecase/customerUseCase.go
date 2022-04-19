package usecase

import (
	"fmt"

	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/SergioVenicio/desafio_nuveo/shared/repositories"
	uuid "github.com/satori/go.uuid"
)

type CustomerUseCase struct {
	Repository repositories.ICustomerRepository
}

func (c *CustomerUseCase) Get(uuid uuid.UUID) (entity.Customer, error) {
	customers, err := c.Repository.Get(uuid)
	if err != nil {
		return entity.Customer{}, err
	}
	return customers, nil
}

func (c *CustomerUseCase) List() []entity.Customer {
	customers, err := c.Repository.ListCustomers()
	if err != nil {
		fmt.Printf("Error on customer list, %s\n", err.Error())
		return nil
	}
	return customers
}

func (c *CustomerUseCase) Create(name string, address string) (entity.Customer, error) {
	newCustomer, err := entity.NewCustomer(name, address)
	if err != nil {
		return entity.Customer{}, err
	}
	err = c.Repository.Create(newCustomer)
	if err != nil {
		return entity.Customer{}, err
	}
	return newCustomer, nil
}

func (c *CustomerUseCase) Delete(uuid uuid.UUID) {
	c.Repository.Delete(uuid)
}

func (c *CustomerUseCase) Update(customer entity.Customer) error {
	err := c.Repository.Update(customer)
	if err != nil {
		return err
	}
	return nil
}
