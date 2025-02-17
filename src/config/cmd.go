package config

import (
	"go-wallet/src/database"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "Simple backend login & register",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
	initCommand.AddCommand(database.SeedCmd)

}

func Run(args []string) error {
	if len(args) == 0 {
		args = append(args, "serve")
	}

	initCommand.SetArgs(args)

	return initCommand.Execute()
}
