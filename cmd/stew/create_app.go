package stew

import (
	"stew/cmd/gofiber"
	"stew/cmd/nodeexpress"
	"stew/cmd/serverless"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var err error
var App = configs.AppConfig{}

// createCmd represents the `create` command.
var createAppCmd = &cobra.Command{
	Use:     "create-service",
	Aliases: []string{"create-app"},
	Short:   "Create a new microservice for your domain in your project's directory",
	Long:    "Create a new microservice for your domain in your project's directory. Always run this from your project directory",
	RunE:    runCreateAppCommand,
}

func runCreateAppCommand(cmd *cobra.Command, args []string) error {
	err = createService()
	showError(err)
	return nil
}

func addServerlessTemplate() error {
	var err error = nil
	serverless.CreateService(Config.ProjectName, App.AppName, App.Runtime, Config.CloudName, Config.Region, App.Environment)
	return err
}

func addTemplate() error {
	var err error = nil
	microserviceTemplate := App.Language + "-" + App.Framework
	switch microserviceTemplate {
	case "go-fiber":
		err = gofiber.CreateMicroservice(App.AppName, App.AppPort)
	case "node-express":
		err = nodeexpress.CreateMicroservice(App.AppName, App.AppPort)
	}

	return err
}

func createBaseService(isContainerBased bool) error {
	// needs to run from project directory
	var template []*survey.Question

	// 1: load project config and project path
	err = Config.LoadConfig()
	showError(err)
	// 2: Get the name of the microservice that needs to be created
	if len(Config.Apps) == 0 {
		// check if its the first app to be created
		err = survey.Ask(surveys.FirstAppQuestion, &App.AppName, survey.WithIcons(surveys.SurveyIconsConfig))
	} else {
		err = survey.Ask(surveys.AppQuestion, &App.AppName, survey.WithIcons(surveys.SurveyIconsConfig))
	}
	showError(err)
	// 3: update projects config
	Config.Apps = append(Config.Apps, App.AppName)
	err = Config.CreateConfig()
	showError(err)
	if isContainerBased {
		// 4: Ask for programming Language
		err = survey.Ask(surveys.LanguageQuestion, &App.Language, survey.WithIcons(surveys.SurveyIconsConfig))
		showError(err)

		// 5: Get Frameworks according to language
		switch App.Language {
		case "go":
			template = surveys.GoQuestions
		case "node":
			template = surveys.NodeQuestions
		}

		// 5: Select Framework
		err = survey.Ask(template, &App.Framework, survey.WithIcons(surveys.SurveyIconsConfig))
		showError(err)
	}

	return nil
}

func createService() error {

	createBaseService(true)
	// ask for app port
	err = survey.Ask(surveys.PortQuestion, &App.AppPort, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	// scaffold the code
	err = addTemplate()
	// add app config file
	err = App.CreateAppConfig()
	showError(err)
	return nil
}

func createServerlessService() error {
	createBaseService(false)
	// ask for region
	err = survey.Ask(surveys.AWSRegion, &Config.Region, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	// ask for cloud provider
	err = survey.Ask(surveys.CloudProviderQuestion, &Config.CloudName, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	err = survey.Ask(surveys.RuntimeQuestion, &App.Runtime, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	err = survey.Ask(surveys.EnvNameQuestion, &App.Environment, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	// scaffold the code
	err = addServerlessTemplate()
	// add app config file
	err = App.CreateAppConfig()
	showError(err)
	return nil
}

func init() {
	rootCmd.AddCommand(createAppCmd)
}
