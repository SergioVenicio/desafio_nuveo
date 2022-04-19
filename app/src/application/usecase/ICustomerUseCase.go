package usecase

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type ICustomerUseCase interface {
	Get(uuid uuid.UUID) (entity.Customer, error)
	List() []entity.Customer
	Create(name string, address string) (entity.Customer, error)
	Update(customer entity.Customer) error
	Delete(uuid uuid.UUID)
	PublishCreateNotification(customer entity.Customer) error
}
