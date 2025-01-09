package config_test

import (
	"testing"

	"github.com/kelvne/psqlmigrator/src/config"
)

var (
	host     string = "localhost"
	port            = "5432"
	dbName          = "testdb"
	user            = "user"
	password        = "password"
)

func TestNewEmpty(t *testing.T) {
	t.Run("TestNewEmpty", func(t *testing.T) {
		cfg := config.NewEmpty()

		if cfg == nil {
			t.Errorf("Expected config to be not nil")
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("TestNew", func(t *testing.T) {
		cfg := config.New(host, port, dbName, user, password)

		if cfg == nil {
			t.Errorf("Expected config to be not nil")
		}
	})
}

func TestConnectionString(t *testing.T) {
	t.Run("TestConnectionString", func(t *testing.T) {
		cfg := config.New(host, port, user, password, dbName)
		expected := "host=localhost port=5432 dbname=testdb user=user password=password sslmode=disable"

		if connString := cfg.ConnectionString(); connString != expected {
			t.Errorf("Expected %s, got %s", expected, connString)
		}
	})
}

func TestRootConnectionString(t *testing.T) {
	t.Run("TestRootConnectionString", func(t *testing.T) {
		cfg := config.New(host, port, user, password, dbName)
		expected := "host=localhost port=5432 user=user password=password sslmode=disable"

		if connString := cfg.RootConnectionString(); connString != expected {
			t.Errorf("Expected %s, got %s", expected, connString)
		}
	})
}
