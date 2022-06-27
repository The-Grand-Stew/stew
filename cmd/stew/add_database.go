package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/cmd/javaspringboot"
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
		err = nodeexpress.AddPostgres()
	case "java-springboot-postgres":
		err = javaspringboot.AddPostgres()
	}
	return err
}

var databaseCmd = &cobra.Command{
	Use:     "add-database",
	Aliases: []string{"db", "database"},
	Short:   "Add code for connecting to a database",
	Long:    "\nThis will add boilerplate code to connect to a particular database inside the microservice. Currently supported databases include postgres,mongo.",
	RunE:    runDbCommand,
}

func createDatabase() error {
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
	commands.ShowMessage("success", fmt.Sprintf("Successfully bootstrapped code in service %s for database %s!", app.AppName, app.Database), true, true)
	//TODO: decide if we want to move to the project directory again here
	return nil

}
func runDbCommand(cmd *cobra.Command, args []string) error {
	return createDatabase()
}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
