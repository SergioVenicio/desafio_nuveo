package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/gorilla/mux"
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

func (c *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	uuid := parameters["uuid"]
	if uuid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("uuid field is required"))
		return
	}

	c.CustomerUseCase.Delete(uuid)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}
