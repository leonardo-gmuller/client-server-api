package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", "data/mydatabase.db")
	if err != nil {
		log.Fatalf("failed to start sqlite: %v", err)
		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("failed to start sqlite: %v", err)
		return nil, err
	}
	err = runSQLFile(DB, "data/init.sql")
	if err != nil {
		log.Fatalf("failed to execute SQL file: %v", err)
		return nil, err
	}
	return DB, nil
}

func runSQLFile(db *sql.DB, filePath string) error {
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return err
	}

	return nil
}
