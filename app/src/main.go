package main

import (
	"github.com/SergioVenicio/desafio_nuveo/application/usecase"
	"github.com/SergioVenicio/desafio_nuveo/shared/http"
	"github.com/SergioVenicio/desafio_nuveo/shared/repositories"
)

func main() {
	useCase := usecase.CustomerUseCase{
		Repository: &repositories.CustomerRepository{},
	}
	config := http.Config{
		Host: "0.0.0.0",
		Port: 5000,
	}
	config.Run(&useCase)
}
