package main

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/pkg/templates"

	"github.com/AlecAivazis/survey/v2"
)

func addTemplate(microserviceTemplate string, appName string) {
	switch microserviceTemplate {
	case "go-fiber":
		gofiber.DownloadTemplate(appName)
	}
}

func addDatabase(databaseTemplate string, appName string) {
	switch databaseTemplate {
	case "go-fiber-postgres":
		gofiber.AddPostgres(appName)
	}
}

func main() {

	var language, framework, database, appName string
	var err error
	var template []*survey.Question
	//Ask for app name
	err = survey.Ask(templates.AppQuestion, &appName, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// Ask Language
	err = survey.Ask(templates.LanguageQuestion, &language, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// Get Frameworks
	switch language {
	case "go":
		template = templates.GoQuestions
	case "python":
		template = templates.PythonQuestions
	}
	survey.Ask(template, &framework, survey.WithIcons(templates.SurveyIconsConfig))
	// Get Database
	survey.Ask(templates.DatabaseQuestions, &database, survey.WithIcons(templates.SurveyIconsConfig))
	// Add Model
	survey.Ask(templates.DatabaseQuestions, &database, survey.WithIcons(templates.SurveyIconsConfig))

	// create a template
	microserviceTemplate := language + "-" + framework
	addTemplate(microserviceTemplate, appName)
	//clone template
	databaseTemplate := microserviceTemplate + "-" + database
	addDatabase(databaseTemplate, appName)
}
