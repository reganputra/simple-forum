package internalsql

import (
	"database/sql"
	"log"
)

func Connect(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}
