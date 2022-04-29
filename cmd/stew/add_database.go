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
	fmt.Printf(databaseTemplate)
	switch databaseTemplate {
	case "go-fiber-postgres":
		err = gofiber.AddPostgres()
	case "node-express-postgres":
		err = nodeexpress.AddPostgres()
	}
	return err
}

var databaseCmd = &cobra.Command{
	Use:     "add-database",
	Aliases: []string{"db", "database"},
	Short:   "Add code to connect to a database",
	Long:    "\nCreate a new db connection via interactive UI.",
	RunE:    runDbCommand,
}

func runDbCommand(cmd *cobra.Command, args []string) error {
	// chec if you are running from the project directory or app directory
	var app configs.AppConfig
	showApplist(app)
	//load the config file
	err := app.LoadAppConfig()
	showError(err)

	commands.ShowMessage("info", fmt.Sprintf("Detected language %s and framework %s for your app", app.Language, app.Framework), true, true)
	// Ask for a database
	err = survey.Ask(surveys.DatabaseQuestions, &app.Database, survey.WithIcons(surveys.SurveyIconsConfig))
	showError(err)
	//Add the database
	template := app.Language + "-" + app.Framework + "-" + app.Database
	err = addDatabase(template, app.AppName)
	if err != nil {
		commands.ShowError(err.Error())
	}
	app.UpdateAppConfig()
	commands.ShowMessage("success", fmt.Sprintf("Successfully bootstrapped code for your database %s!", app.Database), true, true)
	return nil

}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
