package stew

import (
	"fmt"
	"stew/cmd/gofiber"
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
	fmt.Println(app)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Detected Language and framework", app.Language, app.Framework)
	// Ask for a database
	err = survey.Ask(surveys.DatabaseQuestions, &app.Database, survey.WithIcons(surveys.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	//Add the database
	template := app.Language + "-" + app.Framework + "-" + app.Database
	err = addDatabase(template, app.AppName)
	if err != nil {
		fmt.Println(err)
	}
	app.UpdateAppConfig()
	fmt.Println("Success!!")
	return nil

}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
