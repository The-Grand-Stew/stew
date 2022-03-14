package commands

import "github.com/spf13/cobra"

var modelCmd = &cobra.Command{
	Use:  "add model",
	RunE: runModelCmd,
}

func runModelCmd(cmd *cobra.Command, args []string) error {
	// read template name from config file
	// clone or create template for model
	// add files to appropriate foldes
	return nil
}
