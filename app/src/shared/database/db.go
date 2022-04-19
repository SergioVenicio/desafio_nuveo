package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct{}

func (d *Database) OpenConnection() *sql.DB {
	host := os.Getenv("PSQL_HOST")
	user := os.Getenv("PSQL_USER")
	pwd := os.Getenv("PSQL_PWD")
	dbName := os.Getenv("PSQL_DB")

	dbDns := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, pwd, dbName)

	fmt.Println(dbDns)

	db, err := sql.Open("postgres", dbDns)
	if err != nil {
		panic(err.Error())
	}

	return db
}
