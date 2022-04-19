package usecase

import "testing"

func TestCustomerCreateUseCase(t *testing.T) {
	useCase := CustomerUseCase{}

	_, err := useCase.Create("test", "test street 123")
	if err != nil {
		t.Errorf("want nil, get %s", err.Error())
	}
}

func TestCustomerListUseCase(t *testing.T) {
	useCase := CustomerUseCase{}

	useCase.Create("test", "test street 123")
	customers := useCase.List()
	if len(customers) == 0 {
		t.Errorf("want len(customer) > 0, get 0")
	}
}