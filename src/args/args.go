package args

import "github.com/kelvne/psqlmigrator/src/config"

// NewEmpty creates a new empty Args instance, prepared to receive flags.Parse()
func NewEmpty() *Args {
	return &Args{
		config: config.NewEmpty(),
	}
}

// Command returns the command of the Args
func (a *Args) Command() string {
	return a.command
}

// MigrationName returns the migration name to be created
func (a *Args) MigrationName() string {
	return a.migrationName
}

// Config returns the underlying config
func (a *Args) Config() *config.Config {
	return a.config
}
