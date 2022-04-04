package stew

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"stew/cmd/gofiber"
	"stew/cmd/pyfastapi"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var Config = configs.StewConfig{}

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"start"},
	Short:   "Create a new project and add the first service",
	Long:    "Creates a new workspace directory for your project and allow you to add a microservice. Run this command the very first time.",
	RunE:    runCreateCommand,
}

// createCmd represents the `create` command.
var createAppCmd = &cobra.Command{
	Use:     "create-app",
	Aliases: []string{"create-app"},
	Short:   "Create a new microservice in your project directory",
	Long:    "",
	RunE:    runCreateCommand,
}

func addTemplate(microserviceTemplate string, appName string) error {
	var err error = nil
	switch microserviceTemplate {
	case "go-fiber":
		err = gofiber.DownloadTemplate(appName)
	case "python-fastapi":
		err = pyfastapi.DownloadTemplate(appName)
	}
	return err
}

//TODO: Validate if project name is not blank
func createProject() error {
	var err error
	// ask for project name
	err = survey.Ask(surveys.ProjectQuestion, &Config.ProjectName, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		return err
	}
	currentDir, _ := os.Getwd()
	projectPath := filepath.Join(currentDir, Config.ProjectName)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		// create if fir not exists
		err := os.Mkdir(projectPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		// return err that dir exists
		return errors.New(fmt.Sprintf("Directory with the same name as %s exists. Quitting", Config.ProjectName))
	}
	return nil
}

func createService() error {
	var template []*survey.Question
	var err error
	var app = configs.AppDetails{}
	// name of the microservice that needs to be created
	err = survey.Ask(surveys.AppQuestion, &app.AppName, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		return err
	}
	// Ask for programming Language
	err = survey.Ask(surveys.LanguageQuestion, &app.Language, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// Get Frameworks
	switch app.Language {
	case "go":
		template = surveys.GoQuestions
	case "python":
		template = surveys.PythonQuestions
	}
	// get frameworks based on the languages
	err = survey.Ask(template, &app.Framework, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// create a template
	microserviceTemplate := app.Language + "-" + app.Framework
	err = addTemplate(microserviceTemplate, app.AppName)
	if err != nil {
		return err
	}
	Config.Apps = append(Config.Apps, app)
	fmt.Println(Config)
	return nil
}

func runCreateCommand(cmd *cobra.Command, args []string) error {
	err := createProject()
	if err != nil {
		fmt.Println(err)
	}
	currentDir, _ := os.Getwd()
	projectPath := filepath.Join(currentDir, Config.ProjectName)
	os.Chdir(projectPath)
	err = createService()
	if err != nil {
		fmt.Println(err)
	}
	err = Config.CreateConfig()
	if err != nil {
		fmt.Println(err)
	}
	return nil

}

func runCreateAppCommand(cmd *cobra.Command, args []string) error {
	// run from the project directory
	var err error
	// get project name
	currentDir, _ := os.Getwd()
	Config.ProjectName = filepath.Base(currentDir)
	err = createService()
	if err != nil {
		fmt.Println(err)
	}
	err = Config.CreateConfig()
	if err != nil {
		fmt.Println(err)
	}
	return nil

}
func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createAppCmd)
}
