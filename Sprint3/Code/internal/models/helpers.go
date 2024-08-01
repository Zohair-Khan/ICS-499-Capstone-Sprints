package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// This function initializes a database, tests the connection is working.
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
