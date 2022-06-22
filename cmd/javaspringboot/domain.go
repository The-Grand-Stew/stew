package javaspringboot

import (
	"os"
	"path/filepath"
	templates "stew/pkg/templates/javaspringboot"
	"strings"
)

var domainSettings = templates.DomainTemplate{AppName: "", FilePath: ""}

func AddModel(appName string, domain string) error {
	domainSettings.AppName = appName
	err := addModelFile(domain, "pom") // if not present yet
	if err != nil {

		return err
	}
	err = addModelFile(domain, "app") // if not present yet
	if err != nil {

		return err
	}
	err = addModelFile(domain, "controller")
	if err != nil {

		return err
	}
	err = addModelFile(domain, "test")
	if err != nil {

		return err
	}
	return nil
}

func addModelFile(modelName string, fileType string) error {
	currentDir, _ := os.Getwd()
	templateName := ""
	dirPath := ""
	fileName := ""
	domainName := ""
	var method string
	modelNameCamelCase := strings.ToUpper(modelName[:1]) + modelName[1:] // TODO: Discuss how to make camel case
	switch fileType {
	case "pom":
		templateName = templates.JavaSpringBootPOMTemplate
		dirPath = currentDir
		fileName = "pom.xml"
		domainName = modelName
	case "app":
		templateName = templates.JavaSpringBootAppTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", modelName)
		fileName = modelNameCamelCase + "Application.java"
		domainName = modelNameCamelCase
	case "controller":
		templateName = templates.JavaSpringBootControllerTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", modelName)
		fileName = modelNameCamelCase + "Controller.java"
		domainName = modelNameCamelCase
	case "test":
		templateName = templates.JavaSpringBootTestTemplate
		dirPath = filepath.Join(currentDir, "src/test/java/com")
		fileName = modelNameCamelCase + "Test.java"
		domainName = modelNameCamelCase
	}
	domainSettings.FilePath = filepath.Join(dirPath, fileName)
	_ = os.MkdirAll(dirPath, os.ModePerm)     // create directory
	_, _ = os.Create(domainSettings.FilePath) // create file
	domainSettings.TemplateName = templateName
	domainSettings.DomainName = domainName
	domainSettings.Method = method
	// parse
	err := templates.AddSpringBootTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil
}
