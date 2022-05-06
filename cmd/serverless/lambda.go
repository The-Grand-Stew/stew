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
	fmt.Println(lambdaConfig.DirectoryPath)
	if err != nil {
		fmt.Printf("Errored out %s", err)

		return err
	}
	err = addLambdaFile(lambdaConfig.FunctionName, "test", lambdaConfig.Lang)
	if err != nil {
		fmt.Printf("Errored out %s", err)

		return err
	}
	err = addLambdaFile(lambdaConfig.FunctionName, "handler", lambdaConfig.Lang)
	if err != nil {
		fmt.Printf("Errored out handler %s", err)
	}
	err = addLambdaFile(lambdaConfig.FunctionName, "test", lambdaConfig.Lang)
	if err != nil {
		fmt.Printf("Errored out test file %s", err)
	}
	fmt.Println("Updating lambda config")
	functionTemplateString := compileServerlessYamlConfigs(lambdaConfig, templates.ServerlessFunctionConfigYq)
	currentDir, _ := os.Getwd()
	functionFolderPath := filepath.Join(currentDir, "resources", "functions.yml")
	utils.UpdateYmlContents(functionFolderPath, "functions", functionTemplateString)

	return nil
}

func compileServerlessYamlConfigs(templateData templates.LambdaTemplate, templateString string) string {

	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(templateString)
	fmt.Println(t)
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
