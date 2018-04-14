package migration

import (
	"errors"
	"database/sql"

	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/source/file"
)

func MigratePostgresDBUp(db *sql.DB, migrationsDir string) error {
	migrate, err := newPostgresMigration(db, migrationsDir)
	if err != nil {
		return migrationError(err)
	}

	err = migrate.Up()
	if err != nil {
		if err.Error() != "no change" {
			return migrationError(err)
		}
	}
	return nil
}

func newPostgresMigration(db *sql.DB, migrationsDir string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, migrationError(err)
	}
	migrate, err := migrate.NewWithDatabaseInstance("file://" + migrationsDir, "postgres", driver)
	if err != nil {
		return nil, migrationError(err)
	}
	return migrate, nil
}

func migrationError(err error) error {
	return errors.New("migration: " + err.Error())
}