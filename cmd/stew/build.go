package stew

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"build"},
	Short:   "",
	Long:    "",
	RunE:    runBuildCommand,
}

//TODO: Move to buildah,podman
func runBuildCommand(cmd *cobra.Command, args []string) error {
	//load the project config file
	err = Config.LoadConfig()
	showError(err)
	build()
	return nil
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
