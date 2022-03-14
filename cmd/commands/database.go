package commands

import "github.com/spf13/cobra"

var dbCmd = &cobra.Command{
	Use:  "add db",
	RunE: runDatabaseCmd,
}

func runDatabaseCmd(cmd *cobra.Command, args []string) error {
	return nil
}
