package config

import "fmt"

// NewEmpty returns a new reference to an empty Config
func NewEmpty() *Config {
	return new(Config)
}

// New returns a new reference to a Config
func New(host, port, user, password, databaseName string) *Config {
	return &Config{
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
		Username:     user,
		Password:     password,
	}
}

// ConnectionString returns the string for connecting to the given database
func (c *Config) ConnectionString() string {
	return connectionString(c.Host, c.Port, c.Username, c.Password, c.DatabaseName)
}

// RootConectionString returns the string for connecting to the root database
func (c *Config) RootConnectionString() string {
	return connectionString(c.Host, c.Port, c.Username, c.Password, "")
}

func connectionString(host, port, user, password, dbName string) string {
	db := ""
	if len(dbName) > 0 {
		db = "dbname=" + dbName + " "
	}

	return fmt.Sprintf(
		"host=%s port=%s %suser=%s password=%s sslmode=disable",
		host,
		port,
		db,
		user,
		password,
	)
}
