package utils

import (
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, _ := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/sekolah")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
