package stew

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/cmd/tfaws/fargate"
	"stew/pkg/commands"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"
	"strconv"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"deploy"},
	Short:   "Deploy your infrastructure to the cloud with the app you scaffolded",
	Long:    "",
	RunE:    runDeployCommand,
}
var port string

var app configs.AppConfig

func build() string {
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
	// generate tag
	//TODO: HORRIBLE! DO SOMETHING ABOUT THIS
	tag := strconv.FormatInt(time.Now().Unix(), 10)
	commands.ShowMessage("info", fmt.Sprintf("Building image %s:%s", app.AppName, tag), true, true)
	imageName, err := commands.DockerBuild(app.AppName, tag)
	showError(err)
	os.Chdir(currentDir)
	return imageName
}

func deploy(infraPath, imageName string) {
	err = survey.Ask(surveys.EnvNameQuestion, &environment, survey.WithIcons(surveys.SurveyIconsConfig))
	component := Config.CloudName + "-" + Config.CloudComponent
	vars := map[string]string{
		"project":     Config.ProjectName,
		"region":      Config.Region,
		"environment": environment,
		"app_port":    app.AppPort,
		"name":        app.AppName,
		"path_part":   app.AppName,
		"app_image":   "",
	}
	switch component {
	case "aws-ecs-fargate":
		err = fargate.Deploy(infraPath, imageName, vars)
		fmt.Println(err)
		showError(err)
	}
}

func createDeployment() error {
	//load the project config file
	err = Config.LoadConfig()
	showError(err)
	// build the image
	imageName := build()
	commands.ShowMessage("success", "built image "+imageName, true, true)
	currentDir, _ := os.Getwd()
	infraPath := filepath.Join(currentDir, "infrastructure")
	deploy(infraPath, imageName)
	commands.ShowMessage("success", "Infrastructure setup Done!", true, true)
	return nil
}

func runDeployCommand(cmd *cobra.Command, args []string) error {
	return createDeployment()
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
