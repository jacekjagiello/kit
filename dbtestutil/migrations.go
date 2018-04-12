package dbtestutil

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func MigrateDatabaseDown(db *sqlx.DB, migrationsDir string) error {
	migrateTool, err := newMigrationTool(db, migrationsDir)
	if err != nil {
		return err
	}
	return migrateTool.Down()
}

func MigrateDatabaseUp(db *sqlx.DB, migrationsDir string) error {
	migrateTool, err := newMigrationTool(db, migrationsDir)
	if err != nil {
		return err
	}
	return migrateTool.Up()
}

func DropDatabase(db *sqlx.DB, migrationsDir string) error {
	migrateTool, err := newMigrationTool(db, migrationsDir)
	if err != nil {
		return err
	}
	return migrateTool.Drop()
}

func newMigrationTool(db *sqlx.DB, migrationsFolderPath string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migrationTool, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsFolderPath),
		"postgres",
		driver,
	)
	if err != nil {
		return nil, err
	}

	return migrationTool, nil
}