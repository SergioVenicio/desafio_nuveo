package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergioVenicio/desafio_nuveo/application/controllers"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	Host string
	Port int
}

func (hs HttpServer) Run() {
	customerController := controllers.CustomerController{}

	router := mux.NewRouter()
	router.HandleFunc("/cliente", customerController.List).Methods("GET")
	router.HandleFunc("/cliente", customerController.Create).Methods("POST")

	addr := fmt.Sprintf("%s:%d", hs.Host, hs.Port)
	fmt.Println("Running app on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
