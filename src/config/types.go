package config

// Config represents the configuration information for connecting to the database
type Config struct {
	Host         string
	Port         string
	DatabaseName string
	Username     string
	Password     string
}
