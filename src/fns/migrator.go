package fns

import (
	"fmt"

	"github.com/kelvne/psqlmigrator/src/args"
	"github.com/pressly/goose/v3"
)

// NewMigrator returns a new reference to a Migrator
func NewMigrator(arguments *args.Args) *Migrator {
	return &Migrator{
		args: arguments,
	}
}

// Up executes all idle migrations
func (m *Migrator) Up() error {
	if err := m.createDatabase(); err != nil {
		return err
	}

	path, err := migrationsPath()
	if err != nil {
		return err
	}

	return goose.Up(m.sqlDB, path)
}

// Down rollsback the last migration
func (m *Migrator) Down() error {
	if err := m.createDatabase(); err != nil {
		return err
	}

	path, err := migrationsPath()
	if err != nil {
		return err
	}

	return goose.Down(m.sqlDB, path)
}

// Reset resets all migrations
func (m *Migrator) Reset() error {
	if err := m.createDatabase(); err != nil {
		return err
	}

	path, err := migrationsPath()
	if err != nil {
		return err
	}

	return goose.Reset(m.sqlDB, path)
}

func (m *Migrator) connectRoot() (err error) {
	if m.rootDB != nil && m.rootSqlDB != nil {
		if err = m.rootSqlDB.Close(); err != nil {
			return err
		}
	}

	m.rootDB, m.rootSqlDB, err = openGorm(m.args.Config().RootConnectionString())
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrator) connect() (err error) {
	m.db, m.sqlDB, err = openGorm(m.args.Config().ConnectionString())
	if err != nil {
		return err
	}

	return nil
}

func (m *Migrator) isDatabaseCreated() (bool, error) {
	if m.rootDB == nil {
		if err := m.connectRoot(); err != nil {
			return false, err
		}
	}

	result, err := m.rootSqlDB.Exec(fmt.Sprintf(
		"SELECT 1 FROM pg_database WHERE datname='%s'",
		m.args.Config().DatabaseName,
	))
	if err != nil {
		return false, nil
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func (m *Migrator) createDatabase() error {
	exists, err := m.isDatabaseCreated()
	if err != nil {
		return err
	}

	if !exists {
		query := fmt.Sprintf("CREATE DATABASE %s", m.args.Config().DatabaseName)
		if _, err := m.rootSqlDB.Exec(query); err != nil {
			return err
		}
	}

	return m.connect()
}
