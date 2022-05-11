package stew

import (
	"os"
	"os/exec"
	"stew/pkg/commands"
	"stew/pkg/templates/surveys"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:     "play-container",
	Aliases: []string{"play-c"},
	Short:   "CLI tour",
	Long:    "Get a small tour of using the CLI to create and deploy services using stew",
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
	if strings.Trim(answer, " ") == commandStr {
		commands.ShowMessage("command", "\n", true, false)
		return true
	} else {
		commands.ShowError("That doesn't seem like the right command! Try again")
	}
	commands.ShowError("That doesn't seem like the right command!")
	return false
}

const introContent string = `Stew is a CLI tool that creates code scaffolds for developers based on their needs. Just like preparing a stew IRL, stew CLI allows a developer to pick an architecture pattern, add required routes, models etc (ingredients) and finally add some common utility functions for ex: oAuth connectors, custom validation libraries, postgres/mongo/DDB connectors etc (garnishes). Also because we are cloud engineers , we create the infrastructure setup scripts for the application too !. This means that after using stewCLI (once your stew is cooked), a developer has to only add in the bits that matter into the code aka “the business logic”.  Stew is language agnostic and aims to be cloud agnostic (currently supports AWS only) . Yes, that’s right , you can theoretically use stew to scaffold code in ANY language you choose because stew is designed to be use plugins to scaffold code. At this time we have plugins for node js and golang.`

func showIntro() {
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
	} else {
		commands.ShowMessage("doc", "Maybe not today! See you next time", true, true)
		os.Exit(0)
	}

}

var run bool

func runScaffold() {

	commands.ShowMessage("doc", `First, create a new Project called "microservices-demo"`, true, false)
	commands.ShowMessage("doc", `This will also prompt you to create your first service. Let's call it "products" which runs on port "8000". Lets make this service in Go`, true, false)
	if run = generateCommand("stew init"); run {
		createProject()
	}
	commands.ShowMessage("doc", `Great! Let's create another microservice"`, true, true)
	commands.ShowMessage("doc", `First, Go to your project folder`, true, true)
	if run = generateCommand("cd " + Config.ProjectName); run {
		os.Chdir(Config.ProjectName)
	}
	commands.ShowMessage("doc", `Create another service called "payments" that runs on port 8001 and choose node as your preferred language`, true, false)
	if run = generateCommand("stew create-service"); run {
		createService()
	}
	commands.ShowMessage("success", `There! You now have 2 services configured! You can take a look at your project directory to see if you have the boilerplate code generated. (This will open a window in VS code. Make sure you have the command line tool for VS Code installed. Else just open your project in your favorite editor)`, true, true)
	currentDir, _ := os.Getwd()
	var confirm bool
	ready := []*survey.Question{
		{
			Name: "ready",
			Prompt: &survey.Confirm{
				Message: "Want to take a look??",
			},
		},
	}
	survey.Ask(ready, &confirm, survey.WithIcons(surveys.SurveyIconsConfig))
	if confirm {
		commands.ExecCommand("code", []string{currentDir}, true)
	}
	commands.ShowMessage("doc", `Now consider the "products" service uses Postgres database. Lets add some code for connecting to the database. `, true, false)
	if run = generateCommand("stew add-database"); run {
		createDatabase()
		os.Chdir(currentDir)
	}
	commands.ShowMessage("doc", `Next, We realise that the "payments" service needs an additional domain called "internationalPayments". Let's add that in... `, true, false)
	if run = generateCommand("stew add-domain"); run {
		createDomain()
		os.Chdir(currentDir)
	}
	os.Chdir(currentDir)
}

func runDeploy() {
	currentDir, _ := os.Getwd()

	commands.ShowMessage("success", `Now that we have some code in place, let's try deploying to our favorite cloud`, true, false)

	var confirm bool
	ready := []*survey.Question{
		{
			Name: "ready",
			Prompt: &survey.Confirm{
				Message: "Shall we?",
			},
		},
	}
	survey.Ask(ready, &confirm, survey.WithIcons(surveys.SurveyIconsConfig))
	if !confirm {
		commands.ShowMessage("doc", "Maybe not today! See you next time!", true, true)
		os.Exit(0)
	}
	commands.ShowMessage("doc", `The first step is to create a base infrastructure for your microservices. Mostly things like a VPC, API gateways. It depends on the flavor of deployment you choose.`, true, false)
	commands.ShowMessage("doc", `For this example, we are going to deploy to "AWS" using "ECS Fargate" containers. (This might take a while..but hold on!)`, true, false)
	commands.ShowMessage("doc", `Select AWS as your cloud provider with "eu-west-1" as your region and "dev" as your environment when prompted.`, true, false)

	if run = generateCommand("stew setup-infra"); run {
		os.Chdir(currentDir)
		createBaseInfra()
	}

	commands.ShowMessage("sucess", `Good! We have our base setup in place. Time to deploy our services.`, true, false)

	commands.ShowMessage("docs", `Let's deploy the "products" service for now. The same process can be repeated for the other services too"`, true, false)
	if run = generateCommand("stew deploy"); run {
		os.Chdir(currentDir)
		createDeployment()
	}
	commands.ShowMessage("success", "That's All Folks! It was easy, right? You can make your own customisations, write your own business logic or create more services! Its your arena now!", true, true)
}

func runPlayCommand(cmd *cobra.Command, args []string) error {
	showIntro()
	runScaffold()
	runDeploy()
	return nil
}
func init() {
	rootCmd.AddCommand(playCmd)
}
