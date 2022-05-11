package stew

import (
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var playServerlessCmd = &cobra.Command{
	Use:     "play-serverless",
	Aliases: []string{"play-s"},
	Short:   "CLI tour",
	Long:    "Get a small tour of using the CLI to create and deploy services using stew",
	RunE:    runServerlessPlayCommand,
}

func runServerlessScaffold() {

	commands.ShowMessage("doc", `First, create a new Project called "serverless-demo"`, true, false)
	commands.ShowMessage("doc", `This will also prompt you to create your first service. Let's call it "middleware" . Lets make this service in Nodejs`, true, false)
	if run = generateCommand("stew init"); run {
		createProject()
	}
	commands.ShowMessage("success", `There! You now have 1 serverless project configured! You can take a look at your project directory to see if you have the boilerplate code generated. (This will open a window in VS code. Make sure you have the command line tool for VS Code installed. Else just open your project in your favorite editor)`, true, true)
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
	commands.ShowMessage("doc", `Add a lambda function to the "middleware" serverless project. Once you run this command, stew adds the function configuration, handler code and a test file`, true, false)
	if run = generateCommand("stew add-lambda"); run {
		runLambdaCommand()
		os.Chdir(currentDir)
	}
}

func runOffline() {
	currentDir, _ := os.Getwd()

	commands.ShowMessage("success", `Now that we have some code in place, let's try running it offline `, true, false)

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
	appPath := filepath.Join(currentDir, Config.Apps[0])
	commands.ExecCommandWrapper("sls", []string{"offline", "start"}, appPath)
	commands.ShowMessage("success", `Your serverless app is running offline on port 4000`+appPath, true, false)

}

func runServerlessPlayCommand(cmd *cobra.Command, args []string) error {
	showIntro()
	runServerlessScaffold()
	runOffline()
	return nil
}
func init() {
	rootCmd.AddCommand(playServerlessCmd)
}
