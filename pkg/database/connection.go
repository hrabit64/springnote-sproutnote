package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnect() (*sql.DB, error) {

	conn, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		return nil, err
	}

	return conn, err
}
