package serverless

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	templates "stew/pkg/templates/serverless"
	"stew/pkg/utils"
	"strings"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.Title,
}
var lambdaSettings = templates.LambdaTemplate{AppName: "", DirectoryPath: ""}

func AddLambda(appName string, lambdaConfig templates.LambdaTemplate) error {
	lambdaSettings.AppName = appName
	err := addLambdaFile(lambdaConfig.FunctionName, "handler", lambdaConfig.Lang)
	// fmt.Println(lambdaConfig.DirectoryPath)
	if err != nil {
		fmt.Printf("Errored out %s", err)

		return err
	}
	err = addLambdaFile(lambdaConfig.FunctionName, "test", lambdaConfig.Lang)
	if err != nil {
		fmt.Printf("Errored out %s", err)
		return err
	}
	if lambdaConfig.Lang == "nodejs" {
		err = addLambdaFile(lambdaConfig.FunctionName, "packagejson", lambdaConfig.Lang)
		if err != nil {
			fmt.Printf("Errored out %s", err)
			return err
		}
	}
	functionTemplateString := compileServerlessYamlConfigs(lambdaConfig, templates.ServerlessFunctionConfigYq)
	currentDir, _ := os.Getwd()
	// fmt.Println("Attempting to update functions.yml")
	functionFolderPath := filepath.Join(currentDir, "resources", "functions.yml")
	utils.UpdateYmlContents(functionFolderPath, "functions", functionTemplateString)

	return nil
}

func compileServerlessYamlConfigs(templateData templates.LambdaTemplate, templateString string) string {

	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(templateString)
	if err != nil {
		fmt.Println("failed to compile template%w", err)
	}
	var tpl bytes.Buffer
	t.Execute(&tpl, templateData)
	if err != nil {
		fmt.Println("failed to compile serverless template%w", err)
	}
	return tpl.String()
}

func addLambdaFile(lambdaName string, fileType string, language string) error {
	templateName := templates.ServerlessNodeLambda
	switch fileType {
	case "handler":
		if language == "go" {
			templateName = templates.ServerlessGoLambda
		} else if language == "nodejs" {
			templateName = templates.ServerlessNodeLambda
		}

	case "test":
		if language == "go" {
			templateName = templates.ServerlessGoTest
		} else if language == "nodejs" {
			templateName = templates.ServerlessNodeTest
		}
	case "packagejson":
		if language == "nodejs" {
			templateName = templates.ServerlessPackageJSON
		}
	}
	lambdaSettings.TemplateName = templateName
	lambdaSettings.FunctionName = lambdaName
	lambdaSettings.Lang = language
	// parse
	err := templates.AddServerlessTemplate(lambdaSettings, fileType)
	if err != nil {
		return err
	}
	return nil
}
