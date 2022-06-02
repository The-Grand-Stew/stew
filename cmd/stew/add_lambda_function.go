package stew

// import (
// 	"fmt"
// 	"regexp"
// 	serverless "stew/cmd/serverless"
// 	"stew/pkg/configs"
// 	templates "stew/pkg/templates/serverless"
// 	"stew/pkg/templates/surveys"
// 	"stew/pkg/utils"

// 	"github.com/AlecAivazis/survey/v2"
// 	"github.com/spf13/cobra"
// )

// // func addLambda(lang string, appName string) error {
// // 	var err error = nil
// // 	fmt.Printf(lang)
// // 	switch lang {
// // 	case "go":
// // 		err = gofiber.AddPostgres()
// // 	case "node":
// // 		err = nodeexpress.AddPostgres()
// // 	}
// // 	return err
// // }

// var lambdaCmd = &cobra.Command{
// 	Use:     "add-lambda",
// 	Aliases: []string{"lambda"},
// 	Short:   "Add a new lambda function",
// 	Long:    "\nAdd a new lambda function into your existing serverless project",
// 	RunE:    runLambda,
// }

// func runLambda(cmd *cobra.Command, args []string) error {
// 	err := runLambdaCommand()
// 	return err
// }

// func runLambdaCommand() error {
// 	// fmt.Println(app.LoadAppConfig())
// 	//load the config file
// 	var app configs.AppConfig
// 	// Detect the type of config: project or app
// 	showApplist(app)
// 	//load the config
// 	err := app.LoadAppConfig()
// 	showError(err)
// 	logging.ShowMessage("info", fmt.Sprintf("Detected serverless setup with runtime %s", app.Runtime), true, false)

// 	// Ask for a domain name
// 	var lambdaName string
// 	err = survey.Ask(surveys.LambdaNameQuestion, &lambdaName, survey.WithIcons(surveys.SurveyIconsConfig))
// 	showError(err)
// 	// Ask for a domain name
// 	var httpMethod string
// 	err = survey.Ask(surveys.HttpMethodQuestion, &httpMethod, survey.WithIcons(surveys.SurveyIconsConfig))
// 	showError(err)
// 	// Ask for a domain name
// 	var httpPath string
// 	err = survey.Ask(surveys.HttpPathQuestion, &httpPath, survey.WithIcons(surveys.SurveyIconsConfig))
// 	showError(err)
// 	re := regexp.MustCompile("[0-9]+")
// 	// currentDir, _ := os.Getwd()
// 	// lambdaConfig.DirectoryPath =

// 	var lambdaConfig templates.LambdaTemplate
// 	lambdaConfig.AppName = app.AppName
// 	lambdaConfig.FunctionName = lambdaName
// 	lambdaConfig.Environment = app.Environment
// 	lambdaConfig.PathPart = httpPath
// 	lambdaConfig.HttpMethod = httpMethod
// 	langtest := re.Split(app.Runtime, -1)
// 	// fmt.Println("lang", langtest[0])
// 	lambdaConfig.Lang = langtest[0]
// 	extension := utils.ExtensionMap[langtest[0]]
// 	if langtest[0] == "nodejs" {
// 		extension = ".handler"
// 	}
// 	lambdaConfig.HandlerName = "handler" + extension

// 	// add the domain to the code base
// 	err = serverless.AddLambda(app.AppName, lambdaConfig)
// 	showError(err)
// 	app.Lambdas = append(app.Lambdas, lambdaName)
// 	// update config for the app with the additional domain added
// 	err = app.UpdateAppConfig()
// 	showError(err)
// 	logging.ShowMessage("success", fmt.Sprintf("Successfully added Lambda %s!", lambdaName), true, false)
// 	return nil
// }

// func init() {
// 	rootCmd.AddCommand(lambdaCmd)
// }
