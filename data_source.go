package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type dataSource struct {
	DB sqlx.DB
}

func NewDataSource() (*dataSource, error) {
	log.Printf("Initializing Database...\n")

	pgHost := "localhost"
	pgPort := "5432"
	pgUser := "postgres"
	pgPassword := "password"
	pgDatabase := "messaging_backend_db"

	pgConnectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDatabase)

	log.Printf("Connecting to Postgresql\n")
	db, err := sqlx.Open("postgres", pgConnectString)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return &dataSource{
		DB: *db,
	}, nil
}

func (d *dataSource) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing Postgresql: %w", err)
	}

	return nil
}
