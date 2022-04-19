package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/controllers"
	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/gorilla/mux"
)

type Config struct {
	Host string
	Port int
}

func (c *Config) Run(customerUseCase usecase.ICustomerUseCase) {
	customerController := controllers.CustomerController{
		CustomerUseCase: customerUseCase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/cliente", customerController.List).Methods("GET")
	router.HandleFunc("/cliente", customerController.Create).Methods("POST")

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	fmt.Println("Running app on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
