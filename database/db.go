package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Data *sqlx.DB

func ConnectAndMigrate(host, port, databaseName, user, password string, sslMode string) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Printf("ConnectAndMigrate : Error in database connection.")
		return err
	}

	Data = DB
	return migrateUp(DB)
}

func migrateUp(data *sqlx.DB) error {
	driver, err := postgres.WithInstance(data.DB, &postgres.Config{})
	if err != nil {
		log.Printf("migrateUp : Error in retrieving database driver.")
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration",
		"postgres", driver)

	if err != nil {
		log.Printf("migrateUp : Error in creating migrate.")
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("migrateUp : Error in migrating up.")
		return err
	}
	return nil
}
