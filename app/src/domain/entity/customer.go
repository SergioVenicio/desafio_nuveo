package entity

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Customer struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"nome"`
	Address   string    `json:"endereco"`
	CreatedAt string    `json:"cadastrado_em"`
	UpdatedAt string    `json:"atualizado_em"`
}

func NewCustomer(name string, address string) (Customer, error) {
	customer := Customer{
		Uuid:      uuid.NewV4(),
		Name:      name,
		Address:   address,
		CreatedAt: time.Now().UTC().Format("2006-01-02"),
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.Address == "" {
		return errors.New("endereco field is required")
	}

	if c.Name == "" {
		return errors.New("nome field is required")
	}

	return nil
}
