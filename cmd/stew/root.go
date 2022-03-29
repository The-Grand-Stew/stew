package stew

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "stew",
	Version: "0.1",
	Short:   "",
	Long:    ``,
}

func Execute() {
	_ = rootCmd.Execute()
}
