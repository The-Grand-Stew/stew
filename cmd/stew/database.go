package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/pkg/configs"
	"stew/pkg/templates"

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
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Detected Language and framework", cfg.Language, cfg.Framework)
	// Ask for a database
	err = survey.Ask(templates.DatabaseQuestions, &cfg.Database, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	template := cfg.Language + "-" + cfg.Framework + "-" + cfg.Database
	addDatabase(template, cfg.AppName)
	cfg.UpdateConfig()
	fmt.Println("Success!!")
	return nil

}
func init() {
	rootCmd.AddCommand(databaseCmd)
}
