package stew

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var Config = configs.StewConfig{}

// init represents the `init` command.
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"init"},
	Short:   "Create a new project for your services",
	Long:    "Creates a new workspace directory for your project. Run this command the very first time you set up a project consisting of various services. This command will also create the first microservice for you in your project.",
	RunE:    runInitCommand,
}

func showError(err error) {
	if err != nil {
		commands.ShowError(err.Error())
	}
}

func runContainerBased() {
	// create the first Microservice
	err = createService()
	showError(err)
}

func runServerlessBased() {

}

func runInitCommand(cmd *cobra.Command, args []string) error {
	var err error
	// ask for project name
	err = survey.Ask(surveys.ProjectQuestion, &Config.ProjectName, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	// TODO: Ask for project path in options
	currentDir, _ := os.Getwd()
	projectPath := filepath.Join(currentDir, Config.ProjectName)
	err = os.Mkdir(projectPath, os.ModePerm)
	showError(err)
	// change directories to projects path
	os.Chdir(projectPath)

	commands.ShowMessage("success", fmt.Sprintf("Project created at path %s! Go ahead and create your first service", projectPath), false, false)
	// ask for infra type
	err = survey.Ask(surveys.CloudInfraTypeQuestion, &Config.InfrastructureType, survey.WithIcons(surveys.SurveyIconsConfig))
	// change according to infra type
	switch Config.InfrastructureType {
	case "serverless":
		runServerlessBased()
	case "container-based":
		runContainerBased()
	}
	// create a .stew config file in the project directory
	err = Config.CreateConfig()
	showError(err)
	// update the config
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
