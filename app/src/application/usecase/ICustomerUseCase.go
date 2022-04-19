package usecase

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
)

type ICustomerUseCase interface {
	List() []entity.Customer
	Create(name string, address string) (entity.Customer, error)
	Delete(uuid string)
}
