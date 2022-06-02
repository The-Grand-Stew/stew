package serverless

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
	templates "stew/pkg/templates/serverless"
	"stew/pkg/utils"
)

func CreateService(project string, appName string, runtime string, provider string, region string, env string) error {
	gitUrl := repositories.CloudInfraTemplates["serverless-base"]
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	logging.ShowMessage("info", fmt.Sprintf("Cloning Template for serverless application at location : %s", clonePath), true, true)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		logging.ShowMessage("info", fmt.Sprintf("Failed to clone repo : %s", err), true, true)
		return err
	}
	//do npm install
	logging.ShowMessage("info", "Initializing the serverless project", true, true)
	err = commands.NodeInit(clonePath)
	if err != nil {
		logging.ShowMessage("error", fmt.Sprintf("Failed to initialize repo : %s", err), true, true)
		return err
	}
	options := []string{"install", "-g", "husky"}
	command := "npm"
	err = commands.ExecCommandWrapper(command, options, clonePath)
	if err != nil {
		logging.ShowMessage("error", fmt.Sprintf("Failed to install husky: %s", err), true, true)
		return err
	}
	options = []string{"run", "prepare"}
	command = "npm"
	err = commands.ExecCommandWrapper(command, options, clonePath)
	if err != nil {
		logging.ShowMessage("error", fmt.Sprintf("Failed to add git pre-commit hooks: %s", err), true, true)
		return err
	}
	logging.ShowMessage("info", "Updating serverless variables file", true, true)

	varFileName := filepath.Join(clonePath, "variables."+env+".yml")
	varFilePropsTemplate := `{"runtime": "{{ .Runtime }}","environment": "{{ .Environment }}" ,"project": "{{ .Project }}" ,"app": "{{ .AppName }}","region": "{{ .Region }}" }`
	var config templates.LambdaTemplate
	config.AppName = appName
	config.Project = project
	config.Runtime = runtime
	config.Environment = env
	config.Region = region
	vardFileProps := compileServerlessYamlConfigs(config, varFilePropsTemplate)
	utils.UpdateYmlFromRoot(varFileName, ".", vardFileProps, "=")

	pluginName := `serverless-go-build`
	re := regexp.MustCompile("[0-9]+")
	lang := re.Split(runtime, -1)[0]
	if lang == "nodejs" {
		pluginName = `serverless-jest-plugin`
	}

	pluginsFileName := filepath.Join(clonePath, "resources", "plugins.yml")
	utils.UpdateYmlArray(pluginsFileName, "plugins", "\""+pluginName+"\"")
	logging.ShowMessage("info", "Installing required plugin", true, true)
	command = "npm"
	options = []string{"install", "--save-dev", pluginName}
	err = commands.ExecCommandWrapper(command, options, clonePath)
	if err != nil {
		logging.ShowMessage("error", fmt.Sprintf("Failed to install plugin%s", pluginName), true, true)
		return err
	}
	return nil
}
