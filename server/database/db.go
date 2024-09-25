package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", "../data/mydatabase.db")
	if err != nil {
		log.Fatalf("failed to start sqlite: %v", err)
		return nil, err
	}
	return DB, nil
}
