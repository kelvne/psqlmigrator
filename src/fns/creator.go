package fns

import (
	"github.com/pressly/goose/v3"
)

// CreateMigrationDefault creates a new SQL migration on a folder migrations relative to the current working directory
// Eg. $(pwd)/migrations/<new_migration_123.sql>
func CreateMigrationDefault(migrationName string) error {
	path, err := migrationsPath()
	if err != nil {
		return err
	}

	return goose.Create(nil, path, migrationName, "sql")
}
