package repositories

import (
	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/SergioVenicio/desafio_nuveo/shared/database"
)

type CustomerRepository struct {
	Db *database.Database
}

func (r *CustomerRepository) ListCustomers() ([]entity.Customer, error) {
	db := r.Db.OpenConnection()

	statement, err := db.Query("SELECT uuid, name, address, created_at, updated_at FROM customer")
	if err != nil {
		return nil, err
	}

	var customers []entity.Customer
	for statement.Next() {
		var customer entity.Customer

		statement.Scan(&customer.Uuid, &customer.Name, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *CustomerRepository) Create(customer entity.Customer) error {
	db := r.Db.OpenConnection()

	query := "INSERT INTO customer (uuid, name, address, created_at) VALUES ($1, $2, $3, $4)"
	insert, err := db.Prepare(query)
	if err != nil {
		return err
	}

	insert.Exec(customer.Uuid.String(), customer.Name, customer.Address, customer.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
