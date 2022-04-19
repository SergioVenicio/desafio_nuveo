package repositories

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type ICustomerRepository interface {
	Get(uuid uuid.UUID) (entity.Customer, error)
	Update(customer entity.Customer) error
	ListCustomers() ([]entity.Customer, error)
	Create(customer entity.Customer) error
	Delete(uuid uuid.UUID)
}
