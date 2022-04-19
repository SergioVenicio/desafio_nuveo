package entity

import "testing"

func TestNewCustomer(t *testing.T) {
	got, _ := NewCustomer("test", "test street, 123")
	want := Customer{
		Name:    "test",
		Address: "test street, 123",
	}

	if got.Name != want.Name {
		t.Errorf("got name %s, want name %s", want.Name, got.Name)
	}

	if got.Address != want.Address {
		t.Errorf("[]got address %s, want address %s", want.Address, got.Address)
	}
}

func TestValidateCustomer(t *testing.T) {
	customer := Customer{
		Name:    "test",
		Address: "test street, 123",
	}

	err := customer.Validate()
	if err != nil {
		t.Errorf("want nil, get %s", err.Error())
	}
}

func TestValidateCustomerAddressError(t *testing.T) {
	customer := Customer{}

	err := customer.Validate()
	if err == nil {
		t.Errorf("want error, get nil")
	}
}

func TestValidateCustomerNameError(t *testing.T) {
	customer := Customer{
		Address: "test",
	}

	err := customer.Validate()
	if err == nil {
		t.Errorf("want error, get nil")
	}
}

func TestCustomerSetUpdateDate(t *testing.T) {
	customer := Customer{
		Address: "test",
		Name:    "test",
	}
	customer.SetUpdateDate()
	if customer.UpdatedAt == "" {
		t.Errorf("want datetime, get nil")
	}
}
