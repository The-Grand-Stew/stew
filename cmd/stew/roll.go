package commands

import (
	"fmt"
	"stew/pkg/git"
	"stew/pkg/templates"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var rollCmd = &cobra.Command{
	Use:  "roll",
	RunE: runRollCmd,
}

func runRollCmd(cmd *cobra.Command, args []string) error {
	var language, framework, database string
	var err error
	var template []*survey.Question
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
	fmt.Println(framework)
	// Get Database
	survey.Ask(templates.DatabaseQuestions, &database, survey.WithIcons(templates.SurveyIconsConfig))
	fmt.Println(database)
	// create a template
	registry := language + "_" + framework + "_" + database
	//clone template
	git.Clone(registry)
	// add survey
	//create template
	return nil
}
