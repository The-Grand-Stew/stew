package commands

import "github.com/spf13/cobra"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "stew",
	Version: "1.0.0",
	Short:   "",
	Long:    ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	// add roll command
	rootCmd.AddCommand(rollCmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(modelCmd)
}
