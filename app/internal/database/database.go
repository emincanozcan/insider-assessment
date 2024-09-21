package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}

func RunMigrations(databaseURL string) error {
	dbConn, err := NewDB(databaseURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer dbConn.Close()
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func Initialize(databaseURL string) *sql.DB {
	var err error
	err = RunMigrations(databaseURL)
	if err != nil {
		panic("Cant run migrations" + err.Error())
	}

	db, err := NewDB(databaseURL)
	return db
}

