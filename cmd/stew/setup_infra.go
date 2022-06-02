package stew

// import (
// 	"os"
// 	"path/filepath"
// 	"stew/cmd/tfaws"
// 	"stew/cmd/tfaws/fargate"
// 	"stew/pkg/configs"
// 	"stew/pkg/templates/surveys"

// 	"github.com/AlecAivazis/survey/v2"
// 	"github.com/spf13/cobra"
// )

// // init represents the `init` command.
// var setupInfra = &cobra.Command{
// 	Use:     "setup-infrastructure",
// 	Aliases: []string{"setup-infra"},
// 	Short:   "Create base infrastructure for services",
// 	Long:    "Create basic infrastructure components like VPCs, VPN endpoints etc on the chosen cloud provider",
// 	RunE:    setupInfraCommand,
// }

// var InfrastructureConfig configs.InfrastructureConfig
// var environment string

// func createTfStateResources(infraPath string) {
// 	switch Config.CloudName {
// 	case "aws":
// 		vars := map[string]string{"project": Config.ProjectName, "region": Config.Region, "environment": environment}
// 		err = tfaws.SetupTfStateResources(infraPath, vars)
// 		showError(err)
// 	}

// }

// func setCloudComponent() {
// 	err = survey.Ask(surveys.CloudProviderQuestion, &Config.CloudName, survey.WithIcons(surveys.SurveyIconsConfig))
// 	showError(err)
// 	switch Config.CloudName {
// 	case "aws":
// 		// ask for region
// 		err = survey.Ask(surveys.AWSRegion, &Config.Region, survey.WithIcons(surveys.SurveyIconsConfig))
// 		// ask for component name
// 		err = survey.Ask(surveys.AWSComponentQuestion, &Config.CloudComponent, survey.WithIcons(surveys.SurveyIconsConfig))
// 	}
// 	// ask for the env name
// 	err = survey.Ask(surveys.EnvNameQuestion, &environment, survey.WithIcons(surveys.SurveyIconsConfig))
// }

// func runBaseSetup(infraPath string) {
// 	component := Config.CloudName + "-" + Config.CloudComponent
// 	switch component {
// 	case "aws-ecs-fargate":
// 		vars := map[string]string{"project": Config.ProjectName, "region": Config.Region, "environment": environment, "name": "base-setup"}
// 		err = fargate.BaseSetup(infraPath, vars)
// 		showError(err)
// 	}
// }

// func createBaseInfra() error {
// 	// load the project config
// 	err = Config.LoadConfig()
// 	showError(err)
// 	// ask for cloud name
// 	setCloudComponent()
// 	// create directory called infrastructure
// 	currentDir, _ := os.Getwd()
// 	infraPath := filepath.Join(currentDir, "infrastructure")
// 	err = os.Mkdir(infraPath, os.ModePerm)
// 	showError(err)
// 	err = Config.CreateConfig()
// 	// fmt.Println(Config)
// 	// run tf init scripts for setting up tfstate bucket according to cloud
// 	logging.ShowMessage("info", "Creating terraform remote state resources...", true, false)
// 	createTfStateResources(infraPath)
// 	// run base setup according to cloud
// 	logging.ShowMessage("info", "Setting up your Base Infrastructure...", true, true)
// 	// save the config

// 	runBaseSetup(infraPath)
// 	showError(err)
// 	logging.ShowMessage("success", "Base Infrastructure setup done!", true, true)
// 	return nil
// }

// func setupInfraCommand(cmd *cobra.Command, args []string) error {
// 	return createBaseInfra()
// }

// func init() {
// 	rootCmd.AddCommand(setupInfra)
// }
