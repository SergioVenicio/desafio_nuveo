package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type CustomerController struct {
	CustomerUseCase usecase.ICustomerUseCase
}

func (c *CustomerController) List(w http.ResponseWriter, r *http.Request) {
	customers := c.CustomerUseCase.List()
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	err := customer.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	newCustomer, _ := c.CustomerUseCase.Create(customer.Name, customer.Address)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}
