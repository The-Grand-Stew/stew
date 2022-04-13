package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/cmd/nodeexpress"
	"stew/pkg/commands"
	"stew/pkg/configs"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func addDatabase(databaseTemplate string, appName string) error {
	var err error = nil
	switch databaseTemplate {
	case "go-fiber-postgres":
		err = gofiber.AddPostgres()
	case "node-express-postgres":
		err = nodeexpress.AddDatabase(appName, databaseTemplate)
	}
	return err
}

var databaseCmd = &cobra.Command{
	Use:     "add-db",
	Aliases: []string{"db"},
	Short:   "Add code to connect to a database",
	Long:    "\nCreate a new db connection via interactive UI.",
	RunE:    runDbCommand,
}

func runDbCommand(cmd *cobra.Command, args []string) error {
	//load the config file
	var app *configs.AppDetails
	app, err := app.LoadAppConfig()
	if err != nil {
		commands.ShowError(err.Error())
	}
	commands.ShowMessage("info", fmt.Sprintf("Detected language %s and framework %s for your app", app.Language, app.Framework), true, true)
	// Ask for a database
	err = survey.Ask(surveys.DatabaseQuestions, &app.Database, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		commands.ShowError(err.Error())
	}
	//Add the database
	template := app.Language + "-" + app.Framework + "-" + app.Database
	err = addDatabase(template, app.AppName)
	if err != nil {
		commands.ShowError(err.Error())
	}
	app.UpdateAppConfig()
	commands.ShowMessage("success", "Success!", true, true)
	return nil

}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
