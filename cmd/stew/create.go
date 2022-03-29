package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/cmd/pyfastapi"
	"stew/pkg/configs"
	"stew/pkg/templates"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"create"},
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	RunE:    runCreateCommand,
}

func addTemplate(microserviceTemplate string, appName string) error {
	var err error
	switch microserviceTemplate {
	case "go-fiber":
		err = gofiber.DownloadTemplate(appName)
	case "python-fastapi":
		err = pyfastapi.DownloadTemplate(appName)
	}
	return err
}

//TODO: Validate if project name is not blank

func runCreateCommand(cmd *cobra.Command, args []string) error {
	var template []*survey.Question
	var config = configs.StewConfig{}
	var err error
	fmt.Println("Start with your very first project:")
	//Ask for app name
	err = survey.Ask(templates.AppQuestion, &config.AppName, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// Ask Language
	err = survey.Ask(templates.LanguageQuestion, &config.Language, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	// Get Frameworks
	switch config.Language {
	case "go":
		template = templates.GoQuestions
	case "python":
		template = templates.PythonQuestions
	}
	// get frameworks
	survey.Ask(template, &config.Framework, survey.WithIcons(templates.SurveyIconsConfig))
	// create a template
	microserviceTemplate := config.Language + "-" + config.Framework
	fmt.Println("Stewing up your project .....")
	err = addTemplate(microserviceTemplate, config.AppName)
	if err != nil {
		return err
	}
	config.WriteToConfigFile()
	fmt.Println("Success!!")
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
}
