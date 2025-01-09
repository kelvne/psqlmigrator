package main

import (
	"flag"
	"fmt"

	"github.com/kelvne/psqlmigrator/src/args"
	"github.com/kelvne/psqlmigrator/src/fns"
)

var (
	arguments *args.Args
)

func init() {
	arguments = args.ParseFlags()
}

func main() {
	switch arguments.Command() {
	case "create":
		if len(arguments.MigrationName()) == 0 {
			fmt.Println("Migration name is required for creation")
			panic(1)
		}

		if err := fns.CreateMigrationDefault(arguments.MigrationName()); err != nil {
			fmt.Printf("Couldn't create migration, caught error %s", err.Error())
			panic(1)
		}
	case "up", "down", "reset":
		migrator := fns.NewMigrator(arguments)

		fmt.Println("Running migrations...")

		switch arguments.Command() {
		case "up":
			if err := migrator.Up(); err != nil {
				fmt.Printf("Couldn't up migrations, caught error %s", err.Error())
			}
		case "down":
			if err := migrator.Down(); err != nil {
				fmt.Printf("Couldn't down migrations, caught error %s", err.Error())
			}
		case "reset":
			if err := migrator.Reset(); err != nil {
				fmt.Printf("Couldn't reset migrations, caught error %s", err.Error())
			}
		}
	default:
		flag.PrintDefaults()
	}
}
