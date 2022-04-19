package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SergioVenicio/desafio_nuveo/application/controllers"
	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Host            string
	Port            int
	CustomerUseCase usecase.ICustomerUseCase
}

func (c *Config) Run() {
	customerController := controllers.CustomerController{
		CustomerUseCase: c.CustomerUseCase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/cliente/{uuid}", customerController.Delete).Methods("DELETE")
	router.HandleFunc("/cliente", customerController.List).Methods("GET")
	router.HandleFunc("/cliente", customerController.Create).Methods("POST")

	logginRouter := handlers.LoggingHandler(os.Stdout, router)
	jsonRouter := handlers.ContentTypeHandler(logginRouter, "application/json")
	handlers.CompressHandler(jsonRouter)

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	fmt.Println("Running app on", addr)
	log.Fatal(http.ListenAndServe(addr, jsonRouter))
}
