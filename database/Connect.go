package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(connectionString string) (*sql.DB, error) {
	db, error := sql.Open("mysql", connectionString)

	if error != nil {
		return nil, errors.New("Error connecting to a database" + error.Error())
	}

	if error = db.Ping(); error != nil {
		return nil, fmt.Errorf("error connecting to database: %w", error)
	}

	return db, nil
}
