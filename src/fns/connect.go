package fns

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openGorm(connString string) (*gorm.DB, *sql.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(0)
	sqlDB.SetConnMaxLifetime(time.Duration(1) * time.Second)

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, err
	}

	return db, sqlDB, nil
}
