package stew

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/configs"
	"stew/pkg/logging"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
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
		if err == terminal.InterruptErr {
			logging.ShowError("Interrupted!")
			os.Exit(0)
		}
		logging.ShowError(err.Error())

	}
}

func runContainerBased() {
	// create the first Microservice
	err = createService()
	showError(err)
}

// func runServerlessBased() {
// 	// create a serverless project
// 	err = createServerlessService()
// 	showError(err)
// }

func createProject() error {
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

	logging.ShowMessage("success", fmt.Sprintf("Project created at path %s! Go ahead and create your first service....", projectPath), true, true)
	// ask for infra type
	err = survey.Ask(surveys.CloudInfraTypeQuestion, &Config.InfrastructureType, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	// change according to infra type
	// UNCOMMENT ONCE SERVERLESS COMES IN
	switch Config.InfrastructureType {
	// case "serverless":
	// 	runServerlessBased()
	case "container-based":
		runContainerBased()
	}
	// REMOVE AFTER SERVERLESS
	// Config.InfrastructureType = "container-based"
	// runContainerBased()
	// create a .stew config file in the project directory
	err = Config.CreateConfig()
	showError(err)
	// update the config
	return nil
}

func runInitCommand(cmd *cobra.Command, args []string) error {
	return createProject()

}

func init() {
	rootCmd.AddCommand(initCmd)
}
