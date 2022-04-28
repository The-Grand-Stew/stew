package stew

import (
	"errors"
	"fmt"
	"os"
	"stew/cmd/gofiber"
	"stew/cmd/nodeexpress"
	"stew/pkg/commands"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:     "add-domain",
	Aliases: []string{"domain"},
	Short:   "Add another domain to your existing microservice",
	Long:    "",
	RunE:    runDomainCommand,
}

func addDomains(app configs.AppConfig, domain string) error {
	template := app.Language + "-" + app.Framework
	var err error
	switch template {
	case "go-fiber":
		err = gofiber.AddModel(app.AppName, domain)
	case "node-express":
		err = nodeexpress.AddModel(app.AppName, domain)
	}
	return err
}

func showApplist(app configs.AppConfig) {
	configType, appList := configs.DetectConfigType()
	switch configType {
	case 1:
		var appName string
		commands.ShowMessage("info", "Your are running this command from your project directory. You'll have to choose a service", true, false)
		question := surveys.GenerateAppListTemplate(appList)
		//ask for the app to select
		err = survey.Ask(question, &appName, survey.WithIcons(surveys.SurveyIconsConfig))
		app.AppName = appName
	case 0:
		commands.ShowMessage("info", "Your are running this command from your service directory", true, false)
	case -1:
		showError(errors.New("Failed to detect the stew config. Are you in your project directory?"))
	}
	os.Chdir(app.AppName)
}

func runDomainCommand(cmd *cobra.Command, args []string) error {
	//load the config file
	var app configs.AppConfig
	// Detect the type of config: project or app
	showApplist(app)
	//load the config
	err := app.LoadAppConfig()
	showError(err)
	commands.ShowMessage("info", fmt.Sprintf("Detected Language %s and framework %s", app.Language, app.Framework), true, false)

	// Ask for a domain name
	var domain string
	err = survey.Ask(surveys.DomainQuestion, &domain, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)

	// add the domain to the code base
	err = addDomains(app, domain)
	app.Domains = append(app.Domains, domain)
	// update config for the app with the additional domain added
	err = app.UpdateAppConfig()
	showError(err)
	commands.ShowMessage("success", fmt.Sprintf("Successfully created Domain %s!", domain), true, false)
	return nil

}

func init() {
	rootCmd.AddCommand(domainCmd)
}
