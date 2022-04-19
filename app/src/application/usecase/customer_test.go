package usecase

import (
	"testing"

	"github.com/SergioVenicio/desafio_nuveo/shared/repositories"
	uuid "github.com/satori/go.uuid"
)

func TestCustomerCreateUseCase(t *testing.T) {
	useCase := CustomerUseCase{
		Repository: &repositories.MockCustomerRepository{},
	}

	_, err := useCase.Create("test", "test street 123")
	if err != nil {
		t.Errorf("want nil, get %s", err.Error())
	}
}

func TestCustomerListUseCase(t *testing.T) {
	useCase := CustomerUseCase{
		Repository: &repositories.MockCustomerRepository{},
	}

	useCase.Create("test", "test street 123")
	customers := useCase.List()
	if len(customers) == 0 {
		t.Errorf("want len(customer) > 0, get 0")
	}
}

func TestCustomerDeleteUseCase(t *testing.T) {
	useCase := CustomerUseCase{
		Repository: &repositories.MockCustomerRepository{},
	}

	useCase.Create("test", "test street 123")
	customers := useCase.List()
	uuid := customers[0].Uuid

	useCase.Delete(uuid)
	if len(useCase.List()) > 0 {
		t.Errorf("want len(customer) == 0, get %d", len(customers))
	}
}

func TestCustomerGetUseCase(t *testing.T) {
	useCase := CustomerUseCase{
		Repository: &repositories.MockCustomerRepository{},
	}

	customer, _ := useCase.Create("test", "test street 123")
	want, err := useCase.Get(customer.Uuid)
	if err != nil {
		t.Errorf("cant get customer %s", err.Error())
	}
	if want.Uuid != customer.Uuid {
		t.Errorf("want %s, get %s", customer.Uuid, want.Uuid.String())
	}
}

func TestCustomerGetUseCaseWithANonExistentUUID(t *testing.T) {
	useCase := CustomerUseCase{
		Repository: &repositories.MockCustomerRepository{},
	}

	nonExistentUUID := uuid.NewV4()
	_, err := useCase.Get(nonExistentUUID)
	if err == nil {
		t.Errorf("get customer with a non existent UUID %s", nonExistentUUID.String())
	}
}
