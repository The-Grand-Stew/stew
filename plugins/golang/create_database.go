package golang

import (
	"stew/pkg/configs"
	"stew/pkg/logging"
	s "stew/pkg/templates/surveys"
	"stew/plugins/golang/fiber"
	"stew/plugins/golang/surveys"

	"github.com/AlecAivazis/survey/v2"
)

func CreateDatabaseSurvey(App *configs.AppConfig) {
	// ask for framework
	err := survey.Ask(surveys.GoDatabase, &App.Database, survey.WithIcons(s.SurveyIconsConfig))
	if err != nil {
		logging.ShowError(err.Error())
	}
	// get the database name
	if App.Database == "postgres" {
		fiber.AddDatabase(App.Database)
	} else if App.Framework == "echo" {
		logging.ShowMessage("info", "Still in development!!", false, false)
	}
}
