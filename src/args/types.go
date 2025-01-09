package args

import "github.com/kelvne/psqlmigrator/src/config"

// Args represents all the possibe arguments that can be passed to the CLI
type Args struct {
	command       string
	config        *config.Config
	migrationName string
}
