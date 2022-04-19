package main

import (
	"github.com/SergioVenicio/desafio_nuveo/shared/http"
)

func main() {
	server := http.HttpServer{
		Host: "0.0.0.0",
		Port: 5000,
	}
	server.Run()
}
