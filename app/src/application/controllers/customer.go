package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type CustomerController struct {
	useCase usecase.CustomerUseCase
}

func (c *CustomerController) List(w http.ResponseWriter, r *http.Request) {
	customers := c.useCase.List()
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	err := customer.Validate()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	newCustomer, _ := c.useCase.Create(customer.Name, customer.Address)
	json.NewEncoder(w).Encode(newCustomer)
}
