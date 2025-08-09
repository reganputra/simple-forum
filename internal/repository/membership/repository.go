package membership

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {

	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var email string

		err = rows.Scan(&id, &email)
		if err != nil {
			return nil
		}
		log.Printf("User found: %d, %s", id, email)
	}

	return &repository{
		db: db,
	}
}
