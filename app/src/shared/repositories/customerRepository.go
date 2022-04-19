package repositories

import (
	"errors"

	"github.com/SergioVenicio/desafio_nuveo/domain/entity"
	"github.com/SergioVenicio/desafio_nuveo/shared/database"
	uuid "github.com/satori/go.uuid"
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

func (r *CustomerRepository) Get(uuid uuid.UUID) (entity.Customer, error) {
	db := r.Db.OpenConnection()

	var customer entity.Customer

	query := `
		SELECT
			c.uuid
			, c.name
			, c.address
			, COALESCE(TO_CHAR(created_at , 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), '') as created_at
			, COALESCE(TO_CHAR(c.updated_at , 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), '') as updated_at
		FROM customer c WHERE c.uuid = $1`
	row := db.QueryRow(query, uuid)
	err := row.Scan(&customer.Uuid, &customer.Name, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return entity.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (r *CustomerRepository) Update(customer entity.Customer) error {
	db := r.Db.OpenConnection()
	query := `
		UPDATE
			CUSTOMER
		SET
			name = $2
			, address = $3
			, updated_at = $4
		WHERE
			uuid = $1`
	update, err := db.Prepare(query)
	if err != nil {
		return err
	}
	customer.SetUpdateDate()
	_, err = update.Exec(customer.Uuid, customer.Name, customer.Address, customer.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(uuid uuid.UUID) {
	db := r.Db.OpenConnection()

	query := "DELETE FROM customer WHERE uuid = $1"
	delete, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	delete.Exec(uuid)
}
