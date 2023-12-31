package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DataBaseConection() *sql.DB {
	conection := "user=postgres dbname=notebook_store password=12345678 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
