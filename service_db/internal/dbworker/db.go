package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	DB *sql.DB
}

func NewDBConnection(key string) (*sql.DB, error) {
	db, err := sql.Open("postgres", key)
	if err != nil {
		return nil, fmt.Errorf("error occurred while opening the DB: %s", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB: %s", err)
	}

	trans, err := os.ReadFile("init.sql")
	if err != nil {
		return nil, fmt.Errorf("troubles reading init sql script: %s", err)
	}
	_, err = db.Exec(string(trans))
	if err != nil {
		return nil, fmt.Errorf("troubles running init sql script: %s", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping failed: %s", err)
	}

	return db, nil
}
