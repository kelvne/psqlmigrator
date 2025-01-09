package args

import (
	"flag"
)

// ParseFlags parses the command-line flags passed to the CLI and returns an Args struct.
// The following flags are expected:
// -command: Command to be executed <create, up, down, reset>
// -migration-name: Used on <create>, name of the migration
// -host: PostgreSQL host address (default: "localhost")
// -port: PostgreSQL port number (default: "5432")
// -user: Database user (default: "postgres")
// -password: Database password (default: "postgres")
// -dbname: Database name (default: "postgres")
func ParseFlags() *Args {
	args := NewEmpty()

	flag.StringVar(&args.command, "command", "", "Which command to run")
	flag.StringVar(&args.migrationName, "migration-name", "", "Used on <create>, name of the migration")
	flag.StringVar(&args.config.Host, "host", "localhost", "PostgreSQL host address")
	flag.StringVar(&args.config.Port, "port", "5432", "PostgreSQL port number")
	flag.StringVar(&args.config.Username, "user", "postgres", "Database user")
	flag.StringVar(&args.config.Password, "password", "postgres", "Database password")
	flag.StringVar(&args.config.DatabaseName, "dbname", "postgres", "Database name")

	flag.Parse()

	return args
}
