package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // driver for postgres
)

// Database method containing database server connection
func Database() *sql.DB {
	var err error
	connStr := "user=indra dbname=portal sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
