package fns

import (
	"database/sql"

	"github.com/kelvne/psqlmigrator/src/args"
	"gorm.io/gorm"
)

// Migrator represents the instance reponsible for all migration tasks <up, down, reset>
type Migrator struct {
	args      *args.Args
	rootDB    *gorm.DB
	rootSqlDB *sql.DB
	db        *gorm.DB
	sqlDB     *sql.DB
}
