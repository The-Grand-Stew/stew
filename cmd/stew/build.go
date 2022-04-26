package stew

import (
	"stew/pkg/configs"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"build"},
	Short:   "",
	Long:    "",
	RunE:    runBuildCommand,
}

func runBuildCommand(cmd *cobra.Command, args []string) error {
	//load the config file
	var app configs.AppConfig
	// Detect the type of config: project or app
	showApplist(app)
	//load the config
	err := app.LoadAppConfig()
	showError(err)
	return nil

}

func init() {
	rootCmd.AddCommand(buildCmd)
}
