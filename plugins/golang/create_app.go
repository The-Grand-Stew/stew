package golang

import (
	"stew/pkg/configs"
	"stew/pkg/logging"
	s "stew/pkg/templates/surveys"
	"stew/plugins/golang/fiber"
	"stew/plugins/golang/surveys"

	"github.com/AlecAivazis/survey/v2"
)

func CreateAppSurvey(App *configs.AppConfig) {
	// ask for framework
	err := survey.Ask(surveys.GoFramework, &App.Framework, survey.WithIcons(s.SurveyIconsConfig))
	if err != nil {
		logging.ShowError(err.Error())
	}
	//ask for port
	if App.Framework == "fiber" {
		fiber.CreateMicroservice(App.AppName, App.AppPort)
	} else if App.Framework == "echo" {
		logging.ShowMessage("info", "Still in development!!", false, false)
	}
}
