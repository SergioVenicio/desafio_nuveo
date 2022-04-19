package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type CustomerController struct {
	CustomerUseCase usecase.ICustomerUseCase
}

func (c *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	uuid, err := uuid.FromString(parameters["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
		return
	}
	customer, err := c.CustomerUseCase.Get(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
		return
	}
	json.NewEncoder(w).Encode(customer)
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

	err = c.CustomerUseCase.PublishCreateNotification(newCustomer)
	if err != nil {
		panic(err)
	}
}

func (c *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	uuid, err := uuid.FromString(parameters["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("uuid field is required"))
		return
	}

	var customer entity.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	err = customer.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	customer.Uuid = uuid
	err = c.CustomerUseCase.Update(customer)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(""))
}

func (c *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	uuid, err := uuid.FromString(parameters["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("uuid field is required"))
		return
	}

	c.CustomerUseCase.Delete(uuid)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}
