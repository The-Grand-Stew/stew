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
	err := addModelFile(appName, domain, "pom") // if not present yet
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "app") // if not present yet
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "domain")
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "repository")
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "service")
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "controller")
	if err != nil {

		return err
	}
	err = addModelFile(appName, domain, "test") // if not present yet
	if err != nil {

		return err
	}
	return nil
}

func addModelFile(appName string, modelName string, fileType string) error {
	currentDir, _ := os.Getwd()
	templateName := ""
	dirPath := ""
	fileName := ""
	domainName := ""
	var method string
	skipIfFileExists := false
	appNameTitleCase := strings.ToUpper(appName[:1]) + appName[1:]
	modelNameTitleCase := strings.ToUpper(modelName[:1]) + modelName[1:]
	switch fileType {
	case "pom":
		skipIfFileExists = true
		templateName = templates.JavaSpringBootPOMTemplate
		dirPath = currentDir
		fileName = "pom.xml"
		domainName = modelName
	case "app":
		skipIfFileExists = true
		templateName = templates.JavaSpringBootAppTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", appName)
		fileName = appNameTitleCase + "Application.java"
		domainName = modelNameTitleCase
	case "domain":
		templateName = templates.JavaSpringBootDomainTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", appName, "domain")
		fileName = modelNameTitleCase + ".java"
		domainName = modelNameTitleCase
	case "repository":
		templateName = templates.JavaSpringBootRepositoryTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", appName, "repository")
		fileName = modelNameTitleCase + "Repository.java"
		domainName = modelNameTitleCase
	case "service":
		templateName = templates.JavaSpringBootServiceTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", appName, "service")
		fileName = modelNameTitleCase + "Service.java"
		domainName = modelNameTitleCase
	case "controller":
		templateName = templates.JavaSpringBootControllerTemplate
		dirPath = filepath.Join(currentDir, "src/main/java/com", appName, "controller")
		fileName = modelNameTitleCase + "Controller.java"
		domainName = modelNameTitleCase
	case "test":
		skipIfFileExists = true
		templateName = templates.JavaSpringBootTestTemplate
		dirPath = filepath.Join(currentDir, "src/test/java/com", appName, "test")
		fileName = appNameTitleCase + "Test.java"
		domainName = modelNameTitleCase
	}
	domainSettings.FilePath = filepath.Join(dirPath, fileName)
	domainSettings.TemplateName = templateName
	domainSettings.DomainName = domainName
	domainSettings.Method = method
	// parse
	err := templates.AddSpringBootTemplate(domainSettings, skipIfFileExists, dirPath, fileName)
	if err != nil {
		return err
	}
	return nil
}
