package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/pkg/commands"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:     "domain",
	Aliases: []string{"add-domain"},
	Short:   "Add a domain",
	Long:    "\nCreate domains",
	RunE:    runDomainCommand,
}

func addDomains(app *configs.AppDetails) error {
	template := app.Language + "-" + app.Framework
	var err error
	switch template {
	case "go-fiber":
		err = gofiber.AddModel(app.AppName, app.Domains)
		// case "python-fastapi":
		// 	err = pyfastapi.AddModel(app.Domains)
	}
	return err
}

func runDomainCommand(cmd *cobra.Command, args []string) error {
	//load the config file
	var app *configs.AppDetails
	app, err := app.LoadAppConfig()

	if err != nil {
		commands.ShowError(err.Error())
	}
	commands.ShowMessage("info", fmt.Sprintf("Detected Language %s and framework %s", app.Language, app.Framework), true, true)

	// Ask for a database
	var domains string
	err = survey.Ask(surveys.DomainQuestion, &domains, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		commands.ShowError(err.Error())
	}
	domainList := strings.Split(domains, ",")
	app.Domains = append(app.Domains, domainList...)
	addDomains(app)
	app.UpdateAppConfig()
	commands.ShowMessage("success", "Success!", true, true)
	return nil

}

func init() {
	rootCmd.AddCommand(domainCmd)
}
