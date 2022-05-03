package stew

import (
	"os"
	"os/exec"
	"stew/pkg/commands"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// init represents the `init` command.
var playCmd = &cobra.Command{
	Use:     "play",
	Aliases: []string{"play"},
	Short:   "Create a new project for your services",
	Long:    "Creates a new workspace directory for your project. Run this command the very first time you set up a project consisting of various services. This command will also create the first microservice for you in your project.",
	RunE:    runPlayCommand,
}

func generateCommand(commandStr string) bool {
	var answer string
	commands.ShowMessage("command", "Run the command below: "+commandStr, true, true)
	var commandQuestion = []*survey.Question{
		{
			Name:     "cmd",
			Prompt:   &survey.Input{Message: "\n"},
			Validate: survey.Required,
		},
	}
	survey.Ask(commandQuestion, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	if answer == commandStr {
		commands.ShowMessage("command", "\n", true, true)
		return true
	}
	commands.ShowError("That doesn't seem like the right command!")
	return false
}

const advantageContent string = `As cloud engineers/ application developers we work with permutations and combinations of  a set of tools, languages and technologies on our projects. On every project we invest a significant amount of time setting up our codebases, deployment script and pipelines. Some of these tasks can be mundane and at times repitative. A great opportunity to introduce some automation – don’t you think ?​`
const scenarioContent string = ``
const introContent string = `Stew is a CLI tool that creates code scaffolds for developers based on their needs. Just like preparing a stew IRL, stew CLI allows a developer to pick an architecture pattern, add required routes, models etc (ingredients) and finally add some common utility functions for ex: oAuth connectors, custom validation libraries, postgres/mongo/DDB connectors etc (garnishes). Also because we are cloud engineers , we create the infrastructure setup scripts for the application too !. This means that after using stewCLI (once your stew is cooked), a developer has to only add in the bits that matter into the code aka “the business logic”.  Stew is language agnostic and aims to be cloud agnostic (currently supports AWS only) . Yes, that’s right , you can theoretically use stew to scaffold code in ANY language you choose because stew is designed to be use plugins to scaffold code. At this time we have plugins for node js and golang.`
const advantage2Content string = `“Automation is cost cutting by tightening the corners and not cutting them” - This quote describes in a nutshell, what stew aims to be. Stew aims to leverage the expertise and experience of engineers in Deloitte to build high quality code scaffolds ,infrastructure scripts and deployment pipelines for diverse use cases. Using these scaffolds, a development teams reaps the following benefits:​`

func showIntro() bool {
	figure.NewFigure("Stew", "larry3d", true).Print()

	commands.ShowMessage("doc", introContent, true, true)
	var answer string
	next := []*survey.Question{
		{
			Name: "next",
			Prompt: &survey.Input{
				Message: "",
			},
		},
	}
	survey.Ask(next, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	commands.ShowMessage("doc", advantageContent, true, true)
	survey.Ask(next, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	commands.ShowMessage("doc", advantage2Content, true, true)
	survey.Ask(next, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	commands.ShowMessage("doc", "Significant time saving", true, true)
	survey.Ask(next, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	commands.ShowMessage("doc", "Maintain code quality – Readability, maintainability, standardization & documentation​", true, true)
	survey.Ask(next, &answer, survey.WithIcons(surveys.SurveyIconsConfig))
	commands.ShowMessage("doc", "Significant time saving", true, true)
	ready := []*survey.Question{
		{
			Name: "ready",
			Prompt: &survey.Confirm{
				Message: "Let's try it out??",
			},
		},
	}
	var confirm bool
	survey.Ask(ready, &confirm, survey.WithIcons(surveys.SurveyIconsConfig))
	if confirm {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		return true
	}
	commands.ShowMessage("doc", "Maybe not today! See you next time", true, true)
	return false
}

func runPlayCommand(cmd *cobra.Command, args []string) error {
	var run bool
	run = showIntro()
	if !run {
		os.Exit(0)
	}
	commands.ShowMessage("doc", `First, create a new Project called "microservices-demo"`, true, false)
	commands.ShowMessage("doc", `This will also ask you to create your first microservice. Let's call it "productCatalogue" which runs on port 8000`, true, false)
	if run = generateCommand("stew init"); run {
		createProject()
	}
	commands.ShowMessage("doc", `Great! Let's create another microservice"`, true, true)
	commands.ShowMessage("doc", `Create another service called "payments" (Run this from the project directory. But we have done that for you)`, true, false)
	os.Chdir(Config.ProjectName)
	if run = generateCommand("stew create-app"); run {
		createService()
	}
	return nil
}
func init() {
	rootCmd.AddCommand(playCmd)
}
