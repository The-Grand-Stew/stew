package stew

import (
	"os"
	"path/filepath"
	"stew/cmd/tfaws/fargate"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:     "add-domain",
	Aliases: []string{"domain"},
	Short:   "Add another domain to your existing microservice",
	Long:    "",
	RunE:    runDeployCommand,
}
var port string

func getAppPort() string {
	var app configs.AppConfig
	appList := Config.Apps
	question := surveys.GenerateAppListTemplate(appList)
	//ask for the app to select
	var appName string
	err = survey.Ask(question, &appName, survey.WithIcons(surveys.SurveyIconsConfig))
	// load port
	currentDir, _ := os.Getwd()
	os.Chdir(appName)
	err := app.LoadAppConfig()
	showError(err)
	port := app.AppPort
	os.Chdir(currentDir)
	return port

}

func buildImage() {

}

func deploy(infraPath string) {
	component := Config.CloudName + "-" + Config.CloudComponent
	switch component {
	case "aws-ecs-fargate":
		vars := map[string]string{
			"project":     Config.ProjectName,
			"region":      Config.Region,
			"environment": environment,
			"app_port":    port,
		}
		err = fargate.BaseSetup(infraPath, vars)
		showError(err)
	}
}

func runDeployCommand(cmd *cobra.Command, args []string) error {
	//load the project config file
	err = Config.LoadConfig()
	showError(err)
	port = getAppPort()

	currentDir, _ := os.Getwd()
	infraPath := filepath.Join(currentDir, "infrastructure")
	deploy(infraPath)
	return nil

}

func init() {
	rootCmd.AddCommand(deployCmd)
}
